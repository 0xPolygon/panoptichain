package provider

import (
	"context"
	"math"
	"os"
	"testing"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	"github.com/0xPolygon/panoptichain/observer"
)

func newTestClickHouseProvider(t *testing.T, dsn string) *ClickHouseSensorNetworkProvider {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}
	return NewClickHouseSensorNetworkProvider(&network.PolygonMainnet, observer.NewEventBus(),
		config.SensorNetworkClickHouse{Name: "Polygon Mainnet", Label: "wm-test", DSN: dsn})
}

// TestBogusHeadIsIgnored is the regression test for the peer-controlled head.
//
// block_number in block_events is written as announced, so max() alone hands one
// bogon announcement control of the head -- and with it the size of the range the
// fork observer iterates per height (~9,000 years for a 2^62 bogon at ~32ns each).
// The first fix clamped the advance after the fact, which traded the wedge for a
// drifting watermark and, on a first-poll bogon, a permanent silent outage. The
// head query now excludes implausible heights via maxIf, so a mid-life bogon never
// enters the arithmetic at all: the head keeps tracking the REAL maximum while the
// bogon sits in the window, and publishing continues underneath it.
//
// Skipped unless PANOPTICHAIN_TEST_CLICKHOUSE_DSN is set. The bogon row is written
// to a height band no real chain reaches and removed afterwards -- from
// block_events_first as well, because the materialized view fans the insert into
// the rollup and a cleanup that misses it leaves v_block_latency's max(block_number)
// poisoned for the 14-day TTL (found the hard way).
func TestBogusHeadIsIgnored(t *testing.T) {
	dsn := os.Getenv("PANOPTICHAIN_TEST_CLICKHOUSE_DSN")
	if dsn == "" {
		t.Skip("PANOPTICHAIN_TEST_CLICKHOUSE_DSN not set")
	}
	ctx := context.Background()
	p := newTestClickHouseProvider(t, dsn)

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
		VALUES (?, '0xbogonclamptest', 'wm-test', 'deadbeef', 'hash_announce', now(), 0)`, bogon); err != nil {
		t.Fatalf("inject bogon: %v", err)
	}
	t.Cleanup(func() {
		for _, table := range []string{"block_events", "block_events_first"} {
			_ = p.conn.Exec(ctx, `ALTER TABLE `+table+` DELETE WHERE block_hash = '0xbogonclamptest' SETTINGS mutations_sync = 1`)
		}
	})

	if err := p.RefreshState(ctx); err != nil {
		t.Fatalf("second refresh: %v", err)
	}
	// The bogon must be excluded outright: the head keeps tracking the real chain,
	// not drifting toward the bogon a buffer at a time.
	if p.blockNumber >= bogon {
		t.Fatalf("head followed the bogon: %d", p.blockNumber)
	}
	if max := head + blockBufferSize; p.blockNumber > max {
		t.Fatalf("head advanced implausibly with a bogon in the window: head=%d now=%d", head, p.blockNumber)
	}
	// And the watermark keeps advancing normally, so publishing never stopped.
	if p.covered < head {
		t.Fatalf("watermark stalled: covered=%d head at test start=%d", p.covered, head)
	}
	t.Logf("head %d -> %d with a 2^62 bogon in the window (excluded)", head, p.blockNumber)
}

// TestFailedPollRetriesRange is the regression test for the lost-range defect: the
// watermark used to advance before the first query ran, so a poll whose cursor
// failed mid-stream skipped its range forever -- the comments claimed "the next
// poll re-covers the range" and the code did not. The cycle is now transactional:
// nothing commits unless both the blocks and events passes are clean, and the
// watermark only advances on commit, so the next poll retries the same range
// without double-buffering what a failed poll had already read.
func TestFailedPollRetriesRange(t *testing.T) {
	dsn := os.Getenv("PANOPTICHAIN_TEST_CLICKHOUSE_DSN")
	if dsn == "" {
		t.Skip("PANOPTICHAIN_TEST_CLICKHOUSE_DSN not set")
	}
	ctx := context.Background()
	p := newTestClickHouseProvider(t, dsn)

	if err := p.RefreshState(ctx); err != nil {
		t.Fatalf("first refresh: %v", err)
	}
	covered := p.covered
	if covered == 0 {
		t.Skip("no live sensor data")
	}

	// Rewind the watermark so the next fill has a real multi-row range to read,
	// then swap in a connection whose settings make any multi-row result throw
	// mid-cursor -- the same rows.Err() shape a cycle-deadline cancellation
	// produces. The head query (one row) still succeeds, which is the point: the
	// old defect advanced the watermark before the fill ran, so a poll that could
	// see the head but not read the blocks lost the range forever.
	covered -= 20
	p.covered = covered
	limited := NewClickHouseSensorNetworkProvider(&network.PolygonMainnet, observer.NewEventBus(),
		config.SensorNetworkClickHouse{Name: "Polygon Mainnet", Label: "wm-test-limited",
			DSN: dsn + "?max_result_rows=1&result_overflow_mode=throw"})
	if limited.conn == nil {
		t.Fatal("could not build the limited connection")
	}
	goodConn := p.conn
	p.conn = limited.conn
	if err := p.refreshBlockBuffer(ctx); err == nil {
		t.Fatal("expected the fill to fail mid-cursor")
	} else if p.covered != covered {
		t.Fatalf("watermark moved on a failed poll: %d -> %d", covered, p.covered)
	}
	p.conn = goodConn

	// The next clean poll must resume from the untouched watermark and succeed.
	if err := p.refreshBlockBuffer(ctx); err != nil {
		t.Fatalf("recovery refresh: %v", err)
	}
	if p.covered < covered {
		t.Fatalf("watermark went backward: %d -> %d", covered, p.covered)
	}
}

// TestPublishedThroughNeverGoesBackward is the regression test for a false
// sensor_double_sign against a real validator.
//
// Separating the resume watermark from the publish head dropped the property that
// made prevBlockNumber safe: monotonicity. covered legitimately moves DOWN when a
// bogon leaves the head window, and the first version derived prevBlockNumber from
// it -- so the buffer still held blocks above the lowered watermark and the publish
// filter un-skipped them. The same block went out twice in one poll; grouped by
// height with one recovered signer, DoubleSignObserver reads that as a double sign.
//
// No bogon is needed to trigger it: a lagging sensor whose tip announcement ages out
// of the 5-minute window while another announces lower is enough.
//
// publishedThrough is now advanced only through publish(), which cannot lower it.
func TestPublishedThroughNeverGoesBackward(t *testing.T) {
	dsn := os.Getenv("PANOPTICHAIN_TEST_CLICKHOUSE_DSN")
	if dsn == "" {
		t.Skip("PANOPTICHAIN_TEST_CLICKHOUSE_DSN not set")
	}
	p := newTestClickHouseProvider(t, dsn)

	// Drive the state machine directly: the sequence needs a head that goes down,
	// which live data will not do on demand.
	p.covered = 1000
	p.publish(1000)

	// A normal advance.
	p.covered = 1100
	p.publish(1000)
	if p.prevBlockNumber != 1000 {
		t.Fatalf("publish(1000) should set the filter to 1000, got %d", p.prevBlockNumber)
	}
	p.publish(1100)

	// Now the head snaps below the watermark, as it does when a bogon expires.
	p.covered = 1040
	p.publish(1040)
	if p.prevBlockNumber < 1100 {
		t.Fatalf("publish filter went backward to %d; blocks above it are still buffered "+
			"and would republish as a double sign", p.prevBlockNumber)
	}
	if p.publishedThrough != 1100 {
		t.Fatalf("publishedThrough regressed to %d, want 1100", p.publishedThrough)
	}
}

// TestHeadBoundDoesNotWrap is the regression test for the unchecked uint64 addition.
//
// block_number is UInt64 and peer-announced, so a first-poll announcement anywhere in
// the top 512 of the range made covered+blockBufferSize wrap to near zero. The bound
// then sat permanently below the real head, maxIf returned 0 on every poll, and the
// walk was the only recovery path: ~175,000 polls, roughly ten days of silence. This
// is the exact failure class the head bound was introduced to prevent.
func TestHeadBoundDoesNotWrap(t *testing.T) {
	for _, covered := range []uint64{
		math.MaxUint64,
		math.MaxUint64 - 1,
		math.MaxUint64 - blockBufferSize,     // the exact boundary
		math.MaxUint64 - blockBufferSize + 1, // just inside it
		math.MaxUint64 - 2*blockBufferSize,
	} {
		got := addSaturating(covered, blockBufferSize)
		if got < covered {
			t.Fatalf("covered=%d + %d wrapped to %d", covered, blockBufferSize, got)
		}
	}
	// And the ordinary case still adds.
	if got := addSaturating(1000, blockBufferSize); got != 1000+blockBufferSize {
		t.Fatalf("addSaturating(1000, %d) = %d", blockBufferSize, got)
	}
}
