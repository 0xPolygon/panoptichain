package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"testing"
	"time"
)

// usage_requesters is a list of objects nested inside a list of provider
// objects, which is the shape most likely to break silently under viper and
// mapstructure: a decode that drops it leaves usage collection quietly off
// rather than failing to start.
func TestSuccinctProverNetwork_DecodesUsageRequesters(t *testing.T) {
	dir := t.TempDir()
	body := `
namespace: test
providers:
  succinct_prover_network:
    - name: "Succinct Prover Network"
      label: "succinct.xyz"
      url: "rpc.production.succinct.xyz:443"
      api_key: "key"
      usage_requesters:
        - address: "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
          tag: "katana"
        - address: "0xacfe00ba538e753cd0a73ede2c5c27cc44a02fa8"
          tag: "agglayer-node-mainnet"
`
	if err := os.WriteFile(dir+"/config.yml", []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)

	if err := Init(); err != nil {
		t.Fatalf("Init: %v", err)
	}

	networks := Config().Providers.SuccinctProverNetworks
	if len(networks) != 1 {
		t.Fatalf("expected one provider, got %d", len(networks))
	}

	requesters := networks[0].UsageRequesters
	if len(requesters) != 2 {
		t.Fatalf("expected two usage requesters, got %d: %+v", len(requesters), requesters)
	}

	if requesters[0].Address == "" || requesters[0].Tag != "katana" {
		t.Fatalf("requester 0: %+v", requesters[0])
	}

	// The tag is optional, but when given it must survive the decode.
	if requesters[1].Address == "" || requesters[1].Tag != "agglayer-node-mainnet" {
		t.Fatalf("requester 1: %+v", requesters[1])
	}
}

func initSPNConfig(t *testing.T, extra string) error {
	t.Helper()

	dir := t.TempDir()
	body := `
namespace: test
providers:
  succinct_prover_network:
    - name: "Succinct Prover Network"
      label: "succinct.xyz"
      url: "rpc.example.succinct.xyz:443"
      api_key: "key"
` + extra
	if err := os.WriteFile(dir+"/config.yml", []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	certPEM, keyPEM := selfSignedPair(t)
	for name, body := range map[string][]byte{"cert.pem": certPEM, "key.pem": keyPEM} {
		if err := os.WriteFile(dir+"/"+name, body, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	t.Chdir(dir)

	return Init()
}

func TestSuccinctProverNetwork_DecodesClientCertificate(t *testing.T) {
	err := initSPNConfig(t, `
      client_cert: "cert.pem"
      client_key: "key.pem"
`)
	if err != nil {
		t.Fatalf("Init: %v", err)
	}

	n := Config().Providers.SuccinctProverNetworks[0]
	if n.ClientCert == nil || *n.ClientCert != "cert.pem" || n.ClientKey == nil || *n.ClientKey != "key.pem" {
		t.Fatalf("client certificate not decoded: cert=%v key=%v", n.ClientCert, n.ClientKey)
	}
}

// A certificate without its key (or the reverse) would silently connect
// without mTLS, so it must fail at startup instead.
func TestSuccinctProverNetwork_ClientCertificateNeedsBothHalves(t *testing.T) {
	for name, extra := range map[string]string{
		"cert only": "      client_cert: \"cert.pem\"\n",
		"key only":  "      client_key: \"key.pem\"\n",
	} {
		t.Run(name, func(t *testing.T) {
			if err := initSPNConfig(t, extra); err == nil {
				t.Fatal("expected a validation error")
			}
		})
	}
}

func TestSuccinctProverNetwork_ClientCertificateMustExist(t *testing.T) {
	err := initSPNConfig(t, `
      client_cert: "missing.pem"
      client_key: "key.pem"
`)
	if err == nil {
		t.Fatal("expected a validation error for a missing certificate file")
	}
}

// An env reference that expands to nothing leaves both paths as "". That must
// fail at startup rather than pass validation and fail every cycle.
func TestSuccinctProverNetwork_ClientCertificateRejectsEmptyPaths(t *testing.T) {
	err := initSPNConfig(t, `
      client_cert: ""
      client_key: ""
`)
	if err == nil {
		t.Fatal("expected a validation error for empty certificate paths")
	}
}

// selfSignedPair returns a PEM certificate and its matching PEM private key.
func selfSignedPair(t *testing.T) ([]byte, []byte) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "test"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}

	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
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

// Both files exist but are not a matching pair (here, a key for a different
// certificate). That must fail at startup rather than on every cycle.
func TestSuccinctProverNetwork_ClientCertificateMustMatchKey(t *testing.T) {
	_, otherKey := selfSignedPair(t)

	dir := t.TempDir()
	if err := os.WriteFile(dir+"/other-key.pem", otherKey, 0o600); err != nil {
		t.Fatal(err)
	}

	err := initSPNConfig(t, `
      client_cert: "cert.pem"
      client_key: "`+dir+`/other-key.pem"
`)
	if err == nil {
		t.Fatal("expected a startup error for a key that does not match the certificate")
	}
}

func TestSuccinctProverNetwork_ClientCertificateMustBePEM(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(dir+"/garbage.pem", []byte("not a certificate"), 0o600); err != nil {
		t.Fatal(err)
	}

	err := initSPNConfig(t, `
      client_cert: "`+dir+`/garbage.pem"
      client_key: "key.pem"
`)
	if err == nil {
		t.Fatal("expected a startup error for a certificate that is not PEM")
	}
}

// A failed Init must leave the previously loaded config in place rather than a
// partly overwritten one.
func TestInit_FailureKeepsPreviousConfig(t *testing.T) {
	if err := initSPNConfig(t, ""); err != nil {
		t.Fatalf("first Init: %v", err)
	}
	before := Config()

	// Init resolves "." to an absolute search path and keeps it, so remove
	// the first file or the second Init would find it again.
	if err := os.Remove("config.yml"); err != nil {
		t.Fatal(err)
	}

	if err := initSPNConfig(t, "      client_cert: \"cert.pem\"\n"); err == nil {
		t.Fatal("expected the second Init to fail")
	}

	if Config() != before {
		t.Fatal("a failed Init replaced the config")
	}
	if url := Config().Providers.SuccinctProverNetworks[0].URL; url != "rpc.example.succinct.xyz:443" {
		t.Fatalf("config changed after a failed Init: url %q", url)
	}
}
