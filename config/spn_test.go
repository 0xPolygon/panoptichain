package config

import (
	"os"
	"testing"
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
	for _, name := range []string{"cert.pem", "key.pem"} {
		if err := os.WriteFile(dir+"/"+name, []byte("pem"), 0o600); err != nil {
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
