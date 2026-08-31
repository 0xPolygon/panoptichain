package config

import (
	"os"
	"testing"
)

// usage_requesters is a list of objects nested inside a list of provider
// objects, which is the shape most likely to break silently under viper and
// mapstructure: a decode that drops it leaves usage collection quietly off
// rather than failing to start.
func TestSuccinctProverNetwork_DecodesUsageRequestersAndPricing(t *testing.T) {
	dir := t.TempDir()
	body := `
namespace: test
providers:
  succinct_prover_network:
    - name: "Succinct Prover Network"
      label: "succinct.xyz"
      url: "rpc.production.succinct.xyz:443"
      api_key: "key"
      pricing:
        rate_per_billion_gas: 0.5
        currency: "USD"
      usage_requesters:
        - address: "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
          tag: "katana"
          billed: false
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

	pricing := networks[0].Pricing
	if pricing == nil {
		t.Fatal("expected pricing to decode")
	}
	if pricing.RatePerBillionGas != 0.5 || pricing.Currency != "USD" {
		t.Fatalf("pricing = %+v", pricing)
	}

	requesters := networks[0].UsageRequesters
	if len(requesters) != 2 {
		t.Fatalf("expected two usage requesters, got %d: %+v", len(requesters), requesters)
	}

	// An explicit false must survive; the pointer exists so that it can.
	if requesters[0].Tag != "katana" || requesters[0].IsBilled() {
		t.Fatalf("requester 0: tag=%q billed=%v", requesters[0].Tag, requesters[0].IsBilled())
	}

	// An omitted flag means billed, so config only speaks up about exceptions.
	if requesters[1].Tag != "agglayer-node-mainnet" || !requesters[1].IsBilled() {
		t.Fatalf("requester 1: tag=%q billed=%v", requesters[1].Tag, requesters[1].IsBilled())
	}
}
