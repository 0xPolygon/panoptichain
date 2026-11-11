package provider

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/observer"
	spnpb "github.com/0xPolygon/panoptichain/proto/network"
)

type testCA struct {
	cert *x509.Certificate
	key  *ecdsa.PrivateKey
	pool *x509.CertPool
}

func newTestCA(t *testing.T) *testCA {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	tmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "test-ca"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}

	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}

	pool := x509.NewCertPool()
	pool.AddCert(cert)

	return &testCA{cert: cert, key: key, pool: pool}
}

// issue returns a PEM certificate and PEM private key signed by the CA.
func (ca *testCA) issue(t *testing.T, cn string, usage x509.ExtKeyUsage, ips ...net.IP) ([]byte, []byte) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	serial, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		t.Fatal(err)
	}

	tmpl := &x509.Certificate{
		SerialNumber: serial,
		Subject:      pkix.Name{CommonName: cn},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{usage},
		IPAddresses:  ips,
	}

	der, err := x509.CreateCertificate(rand.Reader, tmpl, ca.cert, &key.PublicKey, ca.key)
	if err != nil {
		t.Fatal(err)
	}

	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}

	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}),
		pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})
}

// writeClientCert writes a client key pair into dir, overwriting any previous
// one, and returns the paths.
func writeClientCert(t *testing.T, ca *testCA, dir, cn string) (string, string) {
	t.Helper()

	certPEM, keyPEM := ca.issue(t, cn, x509.ExtKeyUsageClientAuth)
	certPath := filepath.Join(dir, "cert.pem")
	keyPath := filepath.Join(dir, "key.pem")

	if err := os.WriteFile(certPath, certPEM, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keyPath, keyPEM, 0o600); err != nil {
		t.Fatal(err)
	}

	return certPath, keyPath
}

// newMTLSUsageServer starts a gRPC server that requires a client certificate
// signed by ca, answers GetRequesterUsage with one settled hour, and records
// the common name of each client certificate it accepted.
func newMTLSUsageServer(t *testing.T, ca *testCA) (string, func() []string) {
	t.Helper()

	certPEM, keyPEM := ca.issue(t, "server", x509.ExtKeyUsageServerAuth, net.ParseIP("127.0.0.1"))
	serverCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		t.Fatal(err)
	}

	var (
		mu  sync.Mutex
		cns []string
	)

	desc := grpc.ServiceDesc{
		ServiceName: "network.ProverNetwork",
		HandlerType: (*any)(nil),
		Methods: []grpc.MethodDesc{{
			MethodName: "GetRequesterUsage",
			Handler: func(_ any, ctx context.Context, dec func(any) error, _ grpc.UnaryServerInterceptor) (any, error) {
				if err := dec(new(spnpb.GetRequesterUsageRequest)); err != nil {
					return nil, err
				}

				if p, ok := peer.FromContext(ctx); ok {
					if info, ok := p.AuthInfo.(credentials.TLSInfo); ok && len(info.State.PeerCertificates) > 0 {
						mu.Lock()
						cns = append(cns, info.State.PeerCertificates[0].Subject.CommonName)
						mu.Unlock()
					}
				}

				return hourResponse("2026-08-06T12:00:00Z"), nil
			},
		}},
	}

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}

	srv := grpc.NewServer(grpc.Creds(credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    ca.pool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})))
	srv.RegisterService(&desc, new(struct{}))
	go func() { _ = srv.Serve(lis) }()
	t.Cleanup(srv.Stop)

	return lis.Addr().String(), func() []string {
		mu.Lock()
		defer mu.Unlock()
		return append([]string(nil), cns...)
	}
}

func newMTLSProvider(url string, ca *testCA, certPath, keyPath *string) *SuccinctProverNetworkProvider {
	return &SuccinctProverNetworkProvider{
		url:              url,
		logger:           NewLogger(nil, "test"),
		refreshStateTime: new(time.Duration),
		countedHours:     make(map[string]string),
		usageRequesters:  []config.UsageRequester{{Address: "0x5428abf0e5aec1be48597a984a4f9570d9236f29"}},
		rootCAs:          ca.pool,
		clientCert:       certPath,
		clientKey:        keyPath,
	}
}

func refresh(t *testing.T, p *SuccinctProverNetworkProvider) error {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return p.RefreshState(ctx)
}

func TestRefreshState_PresentsClientCertificate(t *testing.T) {
	ca := newTestCA(t)
	url, accepted := newMTLSUsageServer(t, ca)
	certPath, keyPath := writeClientCert(t, ca, t.TempDir(), "panoptichain")

	p := newMTLSProvider(url, ca, &certPath, &keyPath)
	if err := refresh(t, p); err != nil {
		t.Fatalf("RefreshState: %v", err)
	}

	if len(p.usage) != 1 {
		t.Fatalf("expected usage over mTLS, got %d summaries", len(p.usage))
	}

	if got := accepted(); len(got) != 1 || got[0] != "panoptichain" {
		t.Fatalf("server accepted client certificates %v, want [panoptichain]", got)
	}
}

// Without a configured certificate the provider must keep working against a
// server that does not ask for one, and must be rejected by one that does.
func TestRefreshState_NoClientCertificateIsRejected(t *testing.T) {
	ca := newTestCA(t)
	url, accepted := newMTLSUsageServer(t, ca)

	p := newMTLSProvider(url, ca, nil, nil)
	if err := refresh(t, p); err != nil {
		t.Fatalf("RefreshState: %v", err)
	}

	if p.usage != nil {
		t.Fatalf("expected no usage without a client certificate, got %+v", p.usage)
	}

	if got := accepted(); len(got) != 0 {
		t.Fatalf("server accepted client certificates %v, want none", got)
	}
}

// The key pair is re-read every cycle, so replacing the files (as a Kubernetes
// secret update does) rotates the certificate without a restart.
func TestRefreshState_PicksUpRotatedCertificate(t *testing.T) {
	ca := newTestCA(t)
	url, accepted := newMTLSUsageServer(t, ca)
	dir := t.TempDir()
	certPath, keyPath := writeClientCert(t, ca, dir, "before")

	p := newMTLSProvider(url, ca, &certPath, &keyPath)
	if err := refresh(t, p); err != nil {
		t.Fatalf("first RefreshState: %v", err)
	}

	writeClientCert(t, ca, dir, "after")
	if err := refresh(t, p); err != nil {
		t.Fatalf("second RefreshState: %v", err)
	}

	got := accepted()
	if len(got) != 2 || got[0] != "before" || got[1] != "after" {
		t.Fatalf("server accepted client certificates %v, want [before after]", got)
	}
}

// A certificate that cannot be loaded fails the cycle and clears the previous
// cycle's results, so PublishEvents does not republish stale usage.
func TestRefreshState_UnreadableCertificateFailsCycle(t *testing.T) {
	ca := newTestCA(t)
	url, accepted := newMTLSUsageServer(t, ca)
	dir := t.TempDir()
	certPath, keyPath := writeClientCert(t, ca, dir, "panoptichain")

	p := newMTLSProvider(url, ca, &certPath, &keyPath)
	p.usage = []*observer.UsageSummary{{}}

	if err := os.WriteFile(keyPath, []byte("not a key"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := refresh(t, p); err == nil {
		t.Fatal("expected an error for an unreadable key")
	}

	if p.usage != nil {
		t.Fatalf("expected stale usage to be cleared, got %+v", p.usage)
	}

	if got := accepted(); len(got) != 0 {
		t.Fatalf("expected no RPC to reach the server, got %v", got)
	}
}
