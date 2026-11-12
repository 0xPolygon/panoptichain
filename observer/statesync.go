package observer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

// StateSyncCommit is one decoded StateCommitted(uint256 indexed stateId, bool
// success) log from the StateReceiver. It represents the initial contract
// receiver call attempt for a state-sync record: the boolean is the low-level
// CALL result, so success=false means the receiver call failed for any EVM
// reason (revert, out-of-gas under the oracle allowance, ...) — it is NOT
// proof of gas-cap exhaustion. Records delivered to receivers with no code
// emit no event at all and are therefore absent from this denominator.
type StateSyncCommit struct {
	StateId     uint64
	Success     bool
	BlockNumber uint64
	BlockHash   common.Hash
	LogIndex    uint32
}

// StateSyncReplay is one decoded StateSyncReplay(uint256 indexed stateId) log
// from the StateReceiver. The event is only emitted when the replay receiver
// call succeeded (a reverting replay rolls back the event), so every counted
// replay is a canonical successful replay. Replays happen in ordinary
// transactions as well as in Bor's synthetic state-sync receipts. A replay has
// no denominator relationship with commits: an unmatched replay (no observed
// initial failure in retained coverage, e.g. a legacy replay) is still counted.
type StateSyncReplay struct {
	StateId     uint64
	BlockNumber uint64
	BlockHash   common.Hash
	LogIndex    uint32
}

// StateSyncEvents carries the outcome and replay events decoded from one
// refresh cycle's finalized-block scan. It intentionally carries no receiver
// address, payload or error text: those cannot be recovered from the events
// and must not become metric labels.
type StateSyncEvents struct {
	Commits []*StateSyncCommit
	Replays []*StateSyncReplay
}

// StateSyncEventsObserver counts state-sync receiver outcomes and successful
// replays. Both counters are process-local: they reset when the process
// restarts, and each configured provider is counted independently. Multiple
// providers observe the same chain events, so provider series must not be
// summed for chain-level totals.
//
// The supported delivery path publishes each batch at most once. RPCProvider's
// PublishEvents swaps in a fresh batch before EventBus.Publish, and the next
// refresh allocates another batch rather than mutating the published one.
// EventBus invokes each subscriber once per Publish and has no retry or replay
// path; the provider's finalized cursor and seen set handle duplicate RPC logs.
// Notify therefore counts every delivered message without its own dedup state.
// Prometheus CounterVec handles concurrent deliveries from independent
// providers. Directly replaying a message to Notify is outside this contract.
type StateSyncEventsObserver struct {
	commits *prometheus.CounterVec
	replays *prometheus.CounterVec
}

func (o *StateSyncEventsObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)
	events := m.Data().(*StateSyncEvents)
	networkName := m.Network().GetName()
	provider := m.Provider()

	for _, commit := range events.Commits {
		outcome := "failure"
		if commit.Success {
			outcome = "success"
		}

		evt := logger.Debug()
		if !commit.Success {
			// A failed initial receiver call is the rare, interesting case.
			evt = logger.Info()
		}
		evt.
			Uint64("state_id", commit.StateId).
			Uint64("block_number", commit.BlockNumber).
			Str("block_hash", commit.BlockHash.Hex()).
			Uint32("log_index", commit.LogIndex).
			Msg("State sync commit observed")

		o.commits.WithLabelValues(networkName, provider, outcome).Inc()
	}

	for _, replay := range events.Replays {
		logger.Info().
			Uint64("state_id", replay.StateId).
			Uint64("block_number", replay.BlockNumber).
			Str("block_hash", replay.BlockHash.Hex()).
			Uint32("log_index", replay.LogIndex).
			Msg("State sync replay observed")

		o.replays.WithLabelValues(networkName, provider).Inc()
	}
}

func (o *StateSyncEventsObserver) Register(eb *EventBus) {
	eb.Subscribe(topics.BorStateSyncEvents, o)

	o.commits = metrics.NewCounter(
		metrics.RPC,
		"state_sync_commits",
		"The total number of StateCommitted events observed from the StateReceiver, "+
			"i.e. initial contract-receiver state-sync attempts. outcome=success means the "+
			"receiver call succeeded; outcome=failure means it failed for any EVM reason "+
			"(not necessarily out-of-gas). The sum over the outcome label is the number of "+
			"initial contract-receiver attempts; records sent to receivers with no code emit "+
			"no event and are not in this denominator. Each provider is counted independently "+
			"(providers see the same events; do not sum provider series for chain totals). "+
			"Only blocks finalized after the process's initial finalized head are counted; "+
			"there is no startup backfill. Process-local counter: it resets on restart.",
		"outcome",
	)
	o.replays = metrics.NewCounter(
		metrics.RPC,
		"state_sync_replays",
		"The total number of StateSyncReplay events observed from the StateReceiver, "+
			"i.e. canonical successful replays (legacy, hash-backed and historical replays "+
			"included). Replays are a separate recovery signal and are not part of the commit "+
			"attempt denominator; a replay without an observed matching failure is still "+
			"counted. Each provider is counted independently (do not sum provider series for "+
			"chain totals). Only blocks finalized after the process's initial finalized head "+
			"are counted; there is no startup backfill. Process-local counter: it resets on restart.",
	)

	for _, rpc := range config.Config().Providers.RPCs {
		o.commits.WithLabelValues(rpc.Name, rpc.Label, "success").Add(0)
		o.commits.WithLabelValues(rpc.Name, rpc.Label, "failure").Add(0)
		o.replays.WithLabelValues(rpc.Name, rpc.Label).Add(0)
	}
}

func (o *StateSyncEventsObserver) GetCollectors() []prometheus.Collector {
	return []prometheus.Collector{o.commits, o.replays}
}
