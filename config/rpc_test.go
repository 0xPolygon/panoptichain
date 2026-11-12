package config

import (
	"os"
	"testing"
)

// state_sync_events is an explicit per-RPC opt-in. Both an omitted setting and
// an explicit false disable the scan; true must survive the nested RPC decode.
func TestRPC_DecodesStateSyncEvents(t *testing.T) {
	dir := t.TempDir()
	body := `
namespace: test
providers:
  rpc:
    - name: "Polygon Mainnet"
      label: "mainnet.example"
      url: "https://mainnet.example.com"
      state_sync_events: true
      contracts:
        state_sync_receiver_address: "0x0000000000000000000000000000000000001001"
    - name: "Polygon Amoy"
      label: "amoy.example"
      url: "https://amoy.example.com"
      state_sync_events: false
    - name: "Ethereum"
      label: "ethereum.example"
      url: "https://ethereum.example.com"
`
	if err := os.WriteFile(dir+"/config.yml", []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)

	if err := Init(); err != nil {
		t.Fatalf("Init: %v", err)
	}

	rpcs := Config().Providers.RPCs
	if len(rpcs) != 3 {
		t.Fatalf("expected three RPC providers, got %d", len(rpcs))
	}
	if !rpcs[0].StateSyncEvents {
		t.Error("state_sync_events: true decoded as false")
	}
	if rpcs[1].StateSyncEvents {
		t.Error("state_sync_events: false decoded as true")
	}
	if rpcs[2].StateSyncEvents {
		t.Error("omitted state_sync_events did not default to false")
	}
}

func TestRPC_RejectsInvalidStateSyncEvents(t *testing.T) {
	dir := t.TempDir()
	body := `
namespace: test
providers:
  rpc:
    - name: "Polygon Mainnet"
      label: "mainnet.example"
      url: "https://mainnet.example.com"
      state_sync_events: "not-a-boolean"
`
	if err := os.WriteFile(dir+"/config.yml", []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)

	if err := Init(); err == nil {
		t.Fatal("Init accepted a non-boolean state_sync_events value")
	}
}
