package observer

import (
	"bytes"
	"context"
	"os"
	"strings"
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

// stateSyncDeliveryObserver acknowledges each Notify so a test can await the
// real event bus's goroutine delivery.
type stateSyncDeliveryObserver struct {
	*StateSyncEventsObserver
	done chan struct{}
}

func (o *stateSyncDeliveryObserver) Notify(ctx context.Context, m Message) {
	o.StateSyncEventsObserver.Notify(ctx, m)
	o.done <- struct{}{}
}

// Independent providers can publish concurrently through the real bus. Each
// delivery must increment only its own provider series without observer-side
// batch or event dedup state.
func TestStateSyncEvents_ConcurrentIndependentDeliveries(t *testing.T) {
	o := newEventsObserver(t)
	bus := NewEventBus()
	delivery := &stateSyncDeliveryObserver{StateSyncEventsObserver: o, done: make(chan struct{}, 2)}
	bus.Subscribe(topics.BorStateSyncEvents, delivery)

	for _, provider := range []string{"concurrent-a", "concurrent-b"} {
		m := NewMessage(testNetwork, provider, &StateSyncEvents{
			Commits: []*StateSyncCommit{commitEvent(1, false)},
			Replays: []*StateSyncReplay{replayEvent(1)},
		})
		bus.Publish(context.Background(), topics.BorStateSyncEvents, m)
	}
	for range 2 {
		select {
		case <-delivery.done:
		case <-time.After(5 * time.Second):
			t.Fatal("bus delivery timed out")
		}
	}

	for _, provider := range []string{"concurrent-a", "concurrent-b"} {
		if got := counterValue(t, o.commits, testNetwork.GetName(), provider, "failure"); got != 1 {
			t.Errorf("%s failure = %v, want 1", provider, got)
		}
		if got := counterValue(t, o.replays, testNetwork.GetName(), provider); got != 1 {
			t.Errorf("%s replays = %v, want 1", provider, got)
		}
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
