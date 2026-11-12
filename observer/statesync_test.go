package observer

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/0xPolygon/panoptichain/config"
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
		LogIndex:    0,
	}
}

func replayEvent(id uint64) *StateSyncReplay {
	return &StateSyncReplay{
		StateId:     id,
		BlockNumber: 200,
		BlockHash:   common.HexToHash("0xdef"),
		LogIndex:    0,
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
