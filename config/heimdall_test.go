package config

import (
	"os"
	"testing"
	"time"
)

// The fee balance knobs are optional pointers, so a decode that drops them
// silently falls back to the defaults rather than failing to start. Pin the
// shape here: a duration string has to survive viper's decode, and an explicit
// false has to be distinguishable from "unset".
func TestHeimdallEndpoint_DecodesFeeBalanceSettings(t *testing.T) {
	dir := t.TempDir()
	body := `
namespace: test
providers:
  heimdall:
    - name: "Polygon Mainnet"
      label: "polygon.technology"
      tendermint_url: "https://tendermint-api.polygon.technology"
      heimdall_url: "https://heimdall-api.polygon.technology"
      interval: 5s
      fee_balance_interval: 10m
      fee_balance_timeout: 45s
    - name: "Polygon Amoy"
      label: "polygon.technology"
      tendermint_url: "https://tendermint-api-amoy.polygon.technology"
      heimdall_url: "https://heimdall-api-amoy.polygon.technology"
      fee_balances: false
`
	if err := os.WriteFile(dir+"/config.yml", []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)

	if err := Init(); err != nil {
		t.Fatalf("Init: %v", err)
	}

	endpoints := Config().Providers.HeimdallEndpoints
	if len(endpoints) != 2 {
		t.Fatalf("expected two endpoints, got %d", len(endpoints))
	}

	mainnet := endpoints[0]
	if mainnet.FeeBalances != nil {
		t.Errorf("FeeBalances = %v, want nil (unset means enabled)", *mainnet.FeeBalances)
	}
	if mainnet.FeeBalanceInterval == nil || *mainnet.FeeBalanceInterval != 10*time.Minute {
		t.Errorf("FeeBalanceInterval = %v, want 10m", mainnet.FeeBalanceInterval)
	}
	if mainnet.FeeBalanceTimeout == nil || *mainnet.FeeBalanceTimeout != 45*time.Second {
		t.Errorf("FeeBalanceTimeout = %v, want 45s", mainnet.FeeBalanceTimeout)
	}

	amoy := endpoints[1]
	if amoy.FeeBalances == nil || *amoy.FeeBalances {
		t.Errorf("FeeBalances = %v, want an explicit false", amoy.FeeBalances)
	}
	if amoy.FeeBalanceInterval != nil {
		t.Errorf("FeeBalanceInterval = %v, want nil so the default applies", *amoy.FeeBalanceInterval)
	}
}
