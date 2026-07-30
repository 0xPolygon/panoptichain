package provider

import (
	"context"
	"os"
	"testing"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
)

// TestHeadAdvanceIsClamped is the regression test for the peer-controlled head.
//
// block_number in block_events is written as announced, so one bogon announcement
// at an absurd height moves max() -- and with it the End of the range published to
// the fork observer, whose per-height loop then runs once per height of the gap
// (~9,000 years for a 2^63 bogon at ~32ns each). The head may now advance at most
// blockBufferSize per refresh; while a bogon sits in the window the head drifts by
// a buffer per poll instead of jumping, and snaps back when it leaves.
//
// Skipped unless PANOPTICHAIN_TEST_CLICKHOUSE_DSN is set, e.g.
// clickhouse://localhost:9000/sensor against the local-stack database. The bogon
// row is written to a height band no real chain reaches and removed afterwards.
func TestHeadAdvanceIsClamped(t *testing.T) {
	dsn := os.Getenv("PANOPTICHAIN_TEST_CLICKHOUSE_DSN")
	if dsn == "" {
		t.Skip("PANOPTICHAIN_TEST_CLICKHOUSE_DSN not set")
	}
	ctx := context.Background()

	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}

	p := NewClickHouseSensorNetworkProvider(&network.PolygonMainnet, observer.NewEventBus(),
		config.SensorNetworkClickHouse{Name: "Polygon Mainnet", Label: "clamp-test", DSN: dsn})

	// First refresh establishes a real head from live data.
	if err := p.RefreshState(ctx); err != nil {
		t.Fatalf("first refresh: %v", err)
	}
	head := p.blockNumber
	if head == 0 {
		t.Skip("no live sensor data to establish a head")
	}

	// A peer announces a block at an absurd height.
	const bogon = uint64(1) << 62
	if err := p.conn.Exec(ctx, `
		INSERT INTO block_events (block_number, block_hash, sensor_id, node_id, source, seen_at, total_difficulty)
		VALUES (?, '0xbogonclamptest', 'clamp-test', 'deadbeef', 'hash_announce', now(), 0)`, bogon); err != nil {
		t.Fatalf("inject bogon: %v", err)
	}
	t.Cleanup(func() {
		_ = p.conn.Exec(ctx, `ALTER TABLE block_events DELETE WHERE block_hash = '0xbogonclamptest' SETTINGS mutations_sync = 1`)
	})

	if err := p.RefreshState(ctx); err != nil {
		t.Fatalf("second refresh: %v", err)
	}
	if p.blockNumber >= bogon {
		t.Fatalf("head followed the bogon: %d", p.blockNumber)
	}
	if max := head + 2*blockBufferSize; p.blockNumber > max {
		t.Fatalf("head advanced more than a buffer per refresh: head=%d now=%d", head, p.blockNumber)
	}
	t.Logf("head %d -> %d with a 2^62 bogon in the window (clamped)", head, p.blockNumber)
}
