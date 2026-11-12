package observer

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

// Register (and the metrics constructors it calls) dereference the global
// config, so give the test a minimal one.
func newEventsObserver(t *testing.T) *StateSyncEventsObserver {
	t.Helper()

	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}

	o := new(StateSyncEventsObserver)
	o.Register(NewEventBus())

	return o
}

// The event fixtures carry distinct (block hash, log index) identities per
// state ID — as real logs do — so the republication dedup can tell two
// different events apart. Commits and replays live on different hashes, so a
// commit and a replay of the same state ID never collide.
func commitEvent(id uint64, success bool) *StateSyncCommit {
	return &StateSyncCommit{
		StateId:     id,
		Success:     success,
		BlockNumber: 100,
		BlockHash:   common.HexToHash("0xabc"),
		LogIndex:    uint32(id),
	}
}

func replayEvent(id uint64) *StateSyncReplay {
	return &StateSyncReplay{
		StateId:     id,
		BlockNumber: 200,
		BlockHash:   common.HexToHash("0xdef"),
		LogIndex:    uint32(id),
	}
}

// True and false commits land on their outcome series; the sum of the two
// series is the number of initial contract-receiver attempts. The counters
// live on the process-wide registry, so each test uses its own provider
// label to read its series independently.
func TestStateSyncEvents_CountsCommitsByOutcome(t *testing.T) {
	o := newEventsObserver(t)
	provider := "commits"

	o.Notify(context.Background(), NewMessage(testNetwork, provider, &StateSyncEvents{
		Commits: []*StateSyncCommit{
			commitEvent(1, true),
			commitEvent(2, true),
			commitEvent(3, false),
			commitEvent(4, false),
			commitEvent(5, false),
		},
	}))

	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "success"); got != 2 {
		t.Errorf("success = %v, want 2", got)
	}
	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "failure"); got != 3 {
		t.Errorf("failure = %v, want 3", got)
	}
	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "success") +
		counterValue(t, o.commits, testNetwork.GetName(), provider, "failure"); got != 5 {
		t.Errorf("success + failure = %v, want 5 attempts", got)
	}
}

// Replays are counted on their own counter. They are a separate recovery
// signal: a replay without an observed matching failure (an unmatched or
// legacy replay) is still counted, and it never touches the commit series.
func TestStateSyncEvents_CountsReplaysSeparately(t *testing.T) {
	o := newEventsObserver(t)
	provider := "replays"

	o.Notify(context.Background(), NewMessage(testNetwork, provider, &StateSyncEvents{
		Replays: []*StateSyncReplay{
			replayEvent(42),
			replayEvent(43),
		},
	}))

	if got := counterValue(t, o.replays, testNetwork.GetName(), provider); got != 2 {
		t.Errorf("replays = %v, want 2", got)
	}
	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "success"); got != 0 {
		t.Errorf("success = %v, want 0 (replays are not commits)", got)
	}
	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "failure"); got != 0 {
		t.Errorf("failure = %v, want 0 (replays are not commits)", got)
	}
}

// Each provider series counts independently. Both providers observe the same
// chain events; that is exactly why provider series must not be summed for
// chain-level totals.
func TestStateSyncEvents_PerProviderSeries(t *testing.T) {
	o := newEventsObserver(t)
	events := &StateSyncEvents{
		Commits: []*StateSyncCommit{commitEvent(1, false)},
		Replays: []*StateSyncReplay{replayEvent(1)},
	}

	o.Notify(context.Background(), NewMessage(testNetwork, "provider-a", events))
	o.Notify(context.Background(), NewMessage(testNetwork, "provider-b", events))

	if got := counterValue(t, o.commits, testNetwork.GetName(), "provider-a", "failure"); got != 1 {
		t.Errorf("provider-a failure = %v, want 1", got)
	}
	if got := counterValue(t, o.commits, testNetwork.GetName(), "provider-b", "failure"); got != 1 {
		t.Errorf("provider-b failure = %v, want 1", got)
	}
	if got := counterValue(t, o.replays, testNetwork.GetName(), "provider-a"); got != 1 {
		t.Errorf("provider-a replays = %v, want 1", got)
	}
}

// An empty scan publishes nothing; a message with no events changes nothing.
func TestStateSyncEvents_EmptyEvents(t *testing.T) {
	o := newEventsObserver(t)
	provider := "empty"

	o.Notify(context.Background(), NewMessage(testNetwork, provider, &StateSyncEvents{}))

	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "success"); got != 0 {
		t.Errorf("success = %v, want 0", got)
	}
	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "failure"); got != 0 {
		t.Errorf("failure = %v, want 0", got)
	}
	if got := counterValue(t, o.replays, testNetwork.GetName(), provider); got != 0 {
		t.Errorf("replays = %v, want 0", got)
	}
}

// Delivering the same bus message twice — a failed refresh re-publishing a
// stale batch, or any republication — counts each event once per series.
func TestStateSyncEvents_RepublicationCountsOnce(t *testing.T) {
	o := newEventsObserver(t)
	provider := "republication"

	m := NewMessage(testNetwork, provider, &StateSyncEvents{
		Commits: []*StateSyncCommit{commitEvent(1, false)},
		Replays: []*StateSyncReplay{replayEvent(1)},
	})
	o.Notify(context.Background(), m)
	o.Notify(context.Background(), m)

	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "failure"); got != 1 {
		t.Errorf("failure = %v, want 1 (republication deduplicated)", got)
	}
	if got := counterValue(t, o.replays, testNetwork.GetName(), provider); got != 1 {
		t.Errorf("replays = %v, want 1 (republication deduplicated)", got)
	}
}

// The event bus delivers messages from goroutines, so republication can race
// with the original delivery: concurrent delivery of the same message must
// still count each event exactly once.
func TestStateSyncEvents_ConcurrentDeliveryCountsOnce(t *testing.T) {
	o := newEventsObserver(t)
	provider := "concurrent"

	m := NewMessage(testNetwork, provider, &StateSyncEvents{
		Commits: []*StateSyncCommit{commitEvent(1, true)},
		Replays: []*StateSyncReplay{replayEvent(1)},
	})

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.Notify(context.Background(), m)
		}()
	}
	wg.Wait()

	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "success"); got != 1 {
		t.Errorf("success = %v, want 1 under concurrent delivery", got)
	}
	if got := counterValue(t, o.replays, testNetwork.GetName(), provider); got != 1 {
		t.Errorf("replays = %v, want 1 under concurrent delivery", got)
	}
}

// The observed series names, pinned against the actual pinned client_golang
// encoders: both the text and the OpenMetrics exposition call this counter
// test_rpc_state_sync_commits. Neither format appends _total here
// (prometheus/common v0.67.5 does not), and the app's promhttp handler does
// not opt into OpenMetrics — so no _total series is exposed by this app.
func TestStateSyncEvents_ExpositionNames(t *testing.T) {
	o := newEventsObserver(t)
	o.Notify(context.Background(), NewMessage(testNetwork, "exposition",
		&StateSyncEvents{Commits: []*StateSyncCommit{commitEvent(1, true)}}))

	families, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range families {
		if f.GetName() != "test_rpc_state_sync_commits" {
			continue
		}

		var text, om bytes.Buffer
		if _, err := expfmt.MetricFamilyToText(&text, f); err != nil {
			t.Fatal(err)
		}
		if _, err := expfmt.MetricFamilyToOpenMetrics(&om, f); err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(text.String(), "# TYPE test_rpc_state_sync_commits counter") {
			t.Errorf("unexpected text exposition:\n%s", text.String())
		}
		if strings.Contains(om.String(), "state_sync_commits_total") {
			t.Errorf("OpenMetrics exposition unexpectedly appends _total:\n%s", om.String())
		}
		if !strings.Contains(om.String(), "# TYPE test_rpc_state_sync_commits unknown") {
			t.Errorf("unexpected OpenMetrics exposition:\n%s", om.String())
		}
		return
	}

	t.Fatal("state sync commits counter family missing from the registry")
}

// The republication dedup set is bounded like the provider-side scan dedup:
// it keeps the most recent events and evicts the oldest half past the cap.
func TestStateSyncEvents_DedupSetBound(t *testing.T) {
	o := newEventsObserver(t)

	for i := uint64(0); i < stateSyncSeenCap*2; i++ {
		o.markSeen(stateSyncEventKey{
			network:   "n",
			provider:  "p",
			blockHash: common.BigToHash(new(big.Int).SetUint64(i)),
			logIndex:  0,
		})
	}

	o.seenMu.Lock()
	defer o.seenMu.Unlock()
	if len(o.seenKeys) > stateSyncSeenCap {
		t.Errorf("dedup set size = %d, want at most %d", len(o.seenKeys), stateSyncSeenCap)
	}
}

// The recent-batch digest ring is bounded like the per-event set: past the
// cap the oldest half is evicted.
func TestStateSyncEvents_BatchDigestRingBound(t *testing.T) {
	o := newEventsObserver(t)

	for i := uint64(0); i < stateSyncRecentBatchesCap*2; i++ {
		o.claimBatch(stateSyncBatchDigest("n", "p", &StateSyncEvents{
			Commits: []*StateSyncCommit{commitEvent(i, true)},
		}))
	}

	o.batchesMu.Lock()
	defer o.batchesMu.Unlock()
	if len(o.batchDigest) > stateSyncRecentBatchesCap {
		t.Errorf("digest ring size = %d, want at most %d", len(o.batchDigest), stateSyncRecentBatchesCap)
	}
	oldest := stateSyncBatchDigest("n", "p", &StateSyncEvents{
		Commits: []*StateSyncCommit{commitEvent(0, true)},
	})
	if _, ok := o.batchDigest[oldest]; ok {
		t.Error("oldest batch digest retained past the cap")
	}
}

// batchDeliveryObserver acknowledges each Notify so a test can await the
// real event bus's goroutine delivery.
type batchDeliveryObserver struct {
	*StateSyncEventsObserver
	done chan struct{}
}

func (o *batchDeliveryObserver) Notify(ctx context.Context, m Message) {
	o.StateSyncEventsObserver.Notify(ctx, m)
	o.done <- struct{}{}
}

// largeCommitBatch builds a batch of distinct commits spread over several
// blocks, as a multi-chunk scan cycle would accumulate.
func largeCommitBatch(size int) *StateSyncEvents {
	events := &StateSyncEvents{}
	for i := uint64(1); i <= uint64(size); i++ {
		event := commitEvent(i, true)
		event.BlockNumber = 100 + i/100
		event.BlockHash = common.BigToHash(new(big.Int).SetUint64(event.BlockNumber))
		event.LogIndex = uint32(i % 100)
		events.Commits = append(events.Commits, event)
	}
	return events
}

// An immediate republication of a batch larger than the per-event dedup cap
// must still count exactly once per series: whole-batch idempotency is keyed
// on a fixed-size content digest, not on the size-bounded per-event set, so
// FIFO eviction inside the per-event set (which discards the start of such a
// batch during its own delivery) cannot let the same batch count twice. Both
// sides of the per-event cap boundary are covered, through the real event
// bus.
func TestStateSyncEvents_LargeBatchRepublicationCountsOnce(t *testing.T) {
	for _, size := range []int{stateSyncSeenCap - 1, stateSyncSeenCap, stateSyncSeenCap + 1} {
		t.Run(fmt.Sprintf("batch of %d events", size), func(t *testing.T) {
			o := newEventsObserver(t)
			// The counters live on the process-wide registry, so each
			// subtest needs its own provider label.
			provider := fmt.Sprintf("large-batch-%d", size)

			m := NewMessage(testNetwork, provider, largeCommitBatch(size))
			bus := NewEventBus()
			delivery := &batchDeliveryObserver{StateSyncEventsObserver: o, done: make(chan struct{}, 2)}
			bus.Subscribe(topics.BorStateSyncEvents, delivery)

			publishAndWait := func() {
				bus.Publish(context.Background(), topics.BorStateSyncEvents, m)
				select {
				case <-delivery.done:
				case <-time.After(5 * time.Second):
					t.Fatal("bus delivery timed out")
				}
			}

			publishAndWait()
			first := counterValue(t, o.commits, testNetwork.GetName(), provider, "success")
			publishAndWait()
			second := counterValue(t, o.commits, testNetwork.GetName(), provider, "success")

			if first != float64(size) {
				t.Errorf("first delivery counted %g, want %d", first, size)
			}
			if second != first {
				t.Errorf("republication changed the counter: %g -> %g, want unchanged", first, second)
			}
		})
	}
}

// The republication idempotency horizon is bounded by design: the dedup
// state is process-local and size-bounded, so a republication is suppressed
// only while its batch digest is in the recent ring OR its events are in
// the per-event set. This test pushes one batch past both horizons and pins
// the documented behavior: beyond them, a re-delivery counts again.
func TestStateSyncEvents_RepublicationHorizonIsBounded(t *testing.T) {
	o := newEventsObserver(t)
	provider := "horizon"

	first := &StateSyncEvents{Commits: []*StateSyncCommit{commitEvent(1, true)}}
	o.Notify(context.Background(), NewMessage(testNetwork, provider, first))
	if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "success"); got != 1 {
		t.Fatalf("first delivery counted %g, want 1", got)
	}

	// Evict the first batch's event key from the per-event set by counting
	// more than the cap of fresh events...
	o.Notify(context.Background(), NewMessage(testNetwork, provider, largeCommitBatch(stateSyncSeenCap+1)))

	// ...and push its digest out of the recent-batch ring with a stream of
	// distinct small batches.
	for i := uint64(0); i < stateSyncRecentBatchesCap; i++ {
		o.Notify(context.Background(), NewMessage(testNetwork, provider, &StateSyncEvents{
			Commits: []*StateSyncCommit{commitEvent(100_000+i, true)},
		}))
	}

	// Beyond both horizons the same batch counts again — the documented
	// bound of the process-local, size-bounded dedup state.
	before := counterValue(t, o.commits, testNetwork.GetName(), provider, "success")
	o.Notify(context.Background(), NewMessage(testNetwork, provider, first))
	after := counterValue(t, o.commits, testNetwork.GetName(), provider, "success")
	if after != before+1 {
		t.Errorf("republication beyond both dedup horizons did not count again: %g -> %g, want +1", before, after)
	}
}
