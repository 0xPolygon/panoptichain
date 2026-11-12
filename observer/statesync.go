package observer

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/metrics"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

const (
	// stateSyncSeenCap bounds the per-event republication dedup set: the
	// most recently counted distinct event keys are retained, evicting
	// exactly one oldest entry per insertion past the cap (true FIFO, so
	// the last stateSyncSeenCap counted keys are always covered). The
	// store is observer-global across series; like the counters it guards,
	// it is process-local.
	stateSyncSeenCap = 4096

	// stateSyncRecentBatchesCap bounds how many distinct recently
	// completed batches are remembered by digest: the most recently
	// completed digests are retained, evicting exactly one oldest entry
	// per insertion past the cap. A digest is fixed-size regardless of
	// batch size, so whole-batch republication idempotency holds for any
	// supported publication size within this horizon. The store is
	// observer-global across series; like the counters it guards, it is
	// process-local.
	stateSyncRecentBatchesCap = 64
)

// stateSyncEventKey identifies one observed log on one provider series.
// Republication of the same bus message (a failed refresh re-publishing a
// stale batch, or any re-delivery) must not count twice on that series,
// while the same chain event observed via a different provider must still
// count on that provider's own series — hence the provider-scoped key.
type stateSyncEventKey struct {
	network   string
	provider  string
	blockHash common.Hash
	logIndex  uint32
}

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
// Delivery is idempotent per series at two layers:
//
//   - Whole batches are deduplicated by content digest. A claimed digest is
//     pinned as in-flight for the entire active delivery and only enters
//     the bounded completed-batch ring when that delivery finishes. The pin
//     makes the layer eviction-safe under concurrency and provider
//     interleaving: other series' traffic can push a digest out of the
//     bounded completed ring, but never out of the in-flight set while its
//     delivery is still consuming the batch — a concurrent or overlapping
//     republication of an active batch therefore counts nothing. The
//     completed ring retains the stateSyncRecentBatchesCap most recently
//     completed distinct batches. The digest is fixed-size, which makes
//     this layer eviction-safe where a size-bounded per-event set alone is
//     not: during one delivery of a batch larger than the per-event cap,
//     the batch's own earliest keys would already be evicted by its end.
//   - Individual events are deduplicated on (network, provider, block hash,
//     log index), guarding partial overlaps (a later batch re-delivering
//     some already-counted events) within its bounded horizon.
//
// Both layers use bounded, process-local state that resets together with the
// counters on restart, keeping them consistent (a reset counter never skips a
// re-bootstrapped event, and a surviving counter never re-counts one). The
// bounded idempotency horizon is therefore: a republication racing an active
// delivery of the same batch is always suppressed; a republication of a
// completed batch is suppressed while its batch digest is among the
// stateSyncRecentBatchesCap most recently completed distinct batches OR its
// events are among the stateSyncSeenCap most recently counted distinct event
// keys; beyond both horizons a re-delivery counts again. Both stores are
// observer-global across series, so one provider's traffic can consume
// another series' retention horizon — that global scope is intentional and
// bounds total memory. No durable or all-history exactly-once delivery is
// claimed.
type StateSyncEventsObserver struct {
	commits *prometheus.CounterVec
	replays *prometheus.CounterVec

	seenMu    sync.Mutex
	seenKeys  map[stateSyncEventKey]struct{}
	seenOrder []stateSyncEventKey

	batchesMu sync.Mutex
	// inflight pins the digests of deliveries currently being consumed; it
	// is bounded by the number of concurrently active deliveries and
	// drained as each completes.
	inflight map[[sha256.Size]byte]struct{}
	// batchDigest/batchOrder retain the most recently completed distinct
	// batch digests, observer-global, evicting exactly the oldest entry
	// per insertion past the cap.
	batchDigest map[[sha256.Size]byte]struct{}
	batchOrder  [][sha256.Size]byte
}

// stateSyncBatchDigest identifies one delivered batch by content: the series
// it was observed on and every event it carries. Two deliveries of the same
// batch produce the same digest; any difference in series, events or their
// order does not.
func stateSyncBatchDigest(network, provider string, events *StateSyncEvents) [sha256.Size]byte {
	h := sha256.New()
	var word [8]byte

	writeUint64 := func(v uint64) {
		binary.BigEndian.PutUint64(word[:], v)
		h.Write(word[:])
	}

	h.Write([]byte(network))
	h.Write([]byte{0})
	h.Write([]byte(provider))
	h.Write([]byte{0})

	writeUint64(uint64(len(events.Commits)))
	for _, c := range events.Commits {
		h.Write(c.BlockHash[:])
		writeUint64(uint64(c.LogIndex))
		writeUint64(c.StateId)
		writeUint64(c.BlockNumber)
		if c.Success {
			h.Write([]byte{1})
		} else {
			h.Write([]byte{0})
		}
	}

	writeUint64(uint64(len(events.Replays)))
	for _, r := range events.Replays {
		h.Write(r.BlockHash[:])
		writeUint64(uint64(r.LogIndex))
		writeUint64(r.StateId)
		writeUint64(r.BlockNumber)
	}

	var digest [sha256.Size]byte
	h.Sum(digest[:0])
	return digest
}

// claimBatch atomically reports whether this exact batch content is neither
// being delivered right now nor among the recently completed batches,
// pinning the digest as in-flight if so. The pin survives for the whole
// active delivery: other series' concurrent deliveries can evict completed
// digests from the bounded ring, but never an in-flight one, so an
// overlapping republication of a batch whose original delivery is still
// consuming it is always rejected. The event bus delivers messages from
// goroutines, so the check-and-pin must be one critical section. The
// in-flight set is bounded by the number of concurrently active deliveries
// and released when the claiming delivery finishes (completeBatch, called
// via defer — also on panic or early return).
func (o *StateSyncEventsObserver) claimBatch(digest [sha256.Size]byte) bool {
	o.batchesMu.Lock()
	defer o.batchesMu.Unlock()

	if _, ok := o.inflight[digest]; ok {
		return false
	}
	if _, ok := o.batchDigest[digest]; ok {
		return false
	}
	o.inflight[digest] = struct{}{}

	return true
}

// completeBatch releases an in-flight claim and records the digest in the
// bounded completed-batch ring: the most recently completed distinct
// batches are retained (observer-global), evicting exactly the oldest
// entry per insertion past the cap — true FIFO, so a batch still among the
// last stateSyncRecentBatchesCap completed ones is always suppressed. Only
// the delivery that claimed the digest may complete it.
func (o *StateSyncEventsObserver) completeBatch(digest [sha256.Size]byte) {
	o.batchesMu.Lock()
	defer o.batchesMu.Unlock()

	delete(o.inflight, digest)

	if _, ok := o.batchDigest[digest]; ok {
		return
	}
	o.batchDigest[digest] = struct{}{}
	o.batchOrder = append(o.batchOrder, digest)

	if len(o.batchDigest) > stateSyncRecentBatchesCap {
		oldest := o.batchOrder[0]
		delete(o.batchDigest, oldest)
		o.batchOrder = append(o.batchOrder[:0], o.batchOrder[1:]...)
	}
}

// markSeen reports whether the event has not been counted on this
// (network, provider) series yet, recording it if so. The event bus delivers
// messages from goroutines, so check-and-record must be one critical section
// to stay idempotent under concurrent republication. The store is bounded,
// observer-global and a true FIFO: the most recently counted
// stateSyncSeenCap distinct keys are retained, evicting exactly one oldest
// entry per insertion past the cap.
func (o *StateSyncEventsObserver) markSeen(k stateSyncEventKey) bool {
	o.seenMu.Lock()
	defer o.seenMu.Unlock()

	if _, ok := o.seenKeys[k]; ok {
		return false
	}
	o.seenKeys[k] = struct{}{}
	o.seenOrder = append(o.seenOrder, k)

	if len(o.seenKeys) > stateSyncSeenCap {
		oldest := o.seenOrder[0]
		delete(o.seenKeys, oldest)
		o.seenOrder = append(o.seenOrder[:0], o.seenOrder[1:]...)
	}

	return true
}

func (o *StateSyncEventsObserver) Notify(ctx context.Context, m Message) {
	logger := NewLogger(o, m)
	events := m.Data().(*StateSyncEvents)
	networkName := m.Network().GetName()
	provider := m.Provider()

	// Whole-batch idempotency first: a republication of this exact batch
	// content — a failed refresh re-publishing a stale batch, a retry, or a
	// concurrent re-delivery — counts nothing, regardless of batch size.
	// The claim pins the digest as in-flight for this whole delivery, so
	// concurrent republication is rejected even if other series' traffic
	// evicts completed digests from the bounded ring meanwhile; the digest
	// enters the completed ring only when this delivery finishes. The
	// per-event set below still guards partial overlaps within its own
	// horizon.
	digest := stateSyncBatchDigest(networkName, provider, events)
	if !o.claimBatch(digest) {
		return
	}
	defer o.completeBatch(digest)

	series := stateSyncEventKey{network: networkName, provider: provider}

	for _, commit := range events.Commits {
		key := series
		key.blockHash = commit.BlockHash
		key.logIndex = commit.LogIndex
		if !o.markSeen(key) {
			// A republication of an event already counted on this series.
			continue
		}

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
		key := series
		key.blockHash = replay.BlockHash
		key.logIndex = replay.LogIndex
		if !o.markSeen(key) {
			continue
		}

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

	o.seenKeys = make(map[stateSyncEventKey]struct{})
	o.inflight = make(map[[sha256.Size]byte]struct{})
	o.batchDigest = make(map[[sha256.Size]byte]struct{})

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
			"Process-local counter: it resets on restart.",
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
			"chain totals). Process-local counter: it resets on restart.",
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
