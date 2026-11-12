package provider

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/0xPolygon/panoptichain/observer"
)

// Event signatures of the Polygon PoS StateReceiver, computed with
// `cast keccak`:
//
//	StateCommitted(uint256 indexed stateId, bool success)
//	0x5a22725590b0a51c923940223f7458512164b1113359a735e86e7f27f44791ee
//	StateSyncReplay(uint256 indexed stateId)
//	0x8797144948782adcede8e04bfa0bd8fd56941e0df7508bd02a629b477f7b073a
//
// Both events are emitted by the StateReceiver contract (0x…1001 on mainnet
// and Amoy), never by the receiver application: StateCommitted carries only
// the low-level CALL result, so success=false is a receiver failure of any
// EVM kind, not proven gas-cap exhaustion.
var (
	stateCommittedTopic  = common.HexToHash("0x5a22725590b0a51c923940223f7458512164b1113359a735e86e7f27f44791ee")
	stateSyncReplayTopic = common.HexToHash("0x8797144948782adcede8e04bfa0bd8fd56941e0df7508bd02a629b477f7b073a")
)

const (
	// stateSyncScanChunk is the fixed number of blocks requested per
	// eth_getLogs call. It matches the range public Bor RPCs accept.
	stateSyncScanChunk uint64 = 100

	// stateSyncSeenCap bounds the deduplication set. It comfortably covers the
	// bootstrap window plus many cycles of ordinary state-sync traffic.
	stateSyncSeenCap = 4096
)

// stateSyncLogKey uniquely identifies one state-sync log on a chain: the
// block hash it was mined in plus its log index within that block.
type stateSyncLogKey struct {
	blockHash common.Hash
	logIndex  uint32
}

// stateSyncSeenLogs is a bounded FIFO set of log keys that have already been
// counted. It guards against counting the same event twice if an already
// delivered log ever comes back in an overlapping range, a retried response,
// or a republished bus message.
type stateSyncSeenLogs struct {
	keys  map[stateSyncLogKey]struct{}
	order []stateSyncLogKey
}

func newStateSyncSeenLogs() *stateSyncSeenLogs {
	return &stateSyncSeenLogs{keys: make(map[stateSyncLogKey]struct{})}
}

func (s *stateSyncSeenLogs) seen(k stateSyncLogKey) bool {
	_, ok := s.keys[k]
	return ok
}

func (s *stateSyncSeenLogs) add(k stateSyncLogKey) {
	if s.seen(k) {
		return
	}
	s.keys[k] = struct{}{}
	s.order = append(s.order, k)

	if len(s.keys) > stateSyncSeenCap {
		drop := len(s.keys) - stateSyncSeenCap/2
		for _, old := range s.order[:drop] {
			delete(s.keys, old)
		}
		s.order = append(s.order[:0], s.order[drop:]...)
	}
}

// stateSyncLogFetcher fetches the state-sync logs of a bounded block range.
// Splitting it out keeps the cursor, retry and deduplication logic testable
// without a live RPC.
type stateSyncLogFetcher func(ctx context.Context, from, to uint64) ([]types.Log, error)

// refreshStateSyncEvents scans finalized blocks for StateCommitted and
// StateSyncReplay logs when a state sync receiver is configured. It is a
// metrics-only first slice: all scan state is process-local.
func (r *RPCProvider) refreshStateSyncEvents(ctx context.Context, c *ethclient.Client) {
	if r.contracts.StateSyncReceiverAddress == nil {
		return
	}

	receiver := common.HexToAddress(*r.contracts.StateSyncReceiverAddress)

	fetch := func(ctx context.Context, from, to uint64) ([]types.Log, error) {
		query := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(from),
			ToBlock:   new(big.Int).SetUint64(to),
			Addresses: []common.Address{receiver},
			Topics:    [][]common.Hash{{stateCommittedTopic, stateSyncReplayTopic}},
		}
		return c.FilterLogs(ctx, query)
	}

	r.scanStateSyncLogs(ctx, receiver, r.finalizedHeight, fetch)
}

// scanStateSyncLogs advances the provider's state-sync log cursor in fixed
// numeric chunks, accumulating decoded events for this cycle's publication.
// receiver is the configured StateReceiver address; logs emitted by anything
// else are discarded.
//
// Continuity semantics (documented limitations of this first slice):
//
//   - Finality: only blocks up to the current finalized head are scanned, so
//     near-head reorgs cannot be counted twice. Bor emits the state-sync logs
//     in a synthetic receipt with status 1 regardless of the receiver call's
//     outcome, which is why the commit event boolean is counted rather than
//     receipt status. A reorg of already-finalized blocks is a chain integrity
//     failure this collector does not detect.
//   - Cursor: the cursor only advances after a fully successful chunk scan. A
//     failed chunk aborts the scan; the next cycle retries the same range, so
//     blocks are never skipped and successfully scanned chunks are never
//     re-scanned.
//   - Bootstrap: the first scan starts a fixed lookback (blockLookBack, the
//     same default used by the other log filters) below the current finalized
//     head. Coverage starts there, not at genesis: counts are "since coverage
//     start", not all-chain totals.
//   - Deduplication: a bounded (block hash, log index) set prevents the same
//     log from being counted twice across overlapping ranges, retries or
//     republished messages.
//   - Restart: all of this state (cursor, dedup set, counters) is
//     process-local. A restart re-bootstraps and Prometheus counters reset
//     with the process; no durable exactly-once delivery is claimed.
func (r *RPCProvider) scanStateSyncLogs(ctx context.Context, receiver common.Address, finalized uint64, fetch stateSyncLogFetcher) {
	// Each cycle publishes only what this scan accumulates.
	r.stateSyncEvents = &observer.StateSyncEvents{}

	// The finalized head was not fetched this cycle; retry next cycle.
	if finalized == 0 {
		return
	}

	if !r.stateSyncBootstrapped {
		start := uint64(1)
		if finalized > r.blockLookBack+1 {
			start = finalized - r.blockLookBack
		}
		r.stateSyncCursor = start - 1
		r.stateSyncBootstrapped = true

		r.logger.Info().
			Uint64("start_block", start).
			Uint64("finalized_block", finalized).
			Msg("Bootstrapping state sync event scan")
	}

	// A lagging RPC view can report a finalized head below our cursor. Wait
	// for it to catch up rather than rescanning, so counting stays monotonic
	// without skipping blocks.
	if finalized <= r.stateSyncCursor {
		return
	}

	for from := r.stateSyncCursor + 1; from <= finalized; from += stateSyncScanChunk {
		to := min(from+stateSyncScanChunk-1, finalized)

		logs, err := fetch(ctx, from, to)
		if err != nil {
			// Leave the cursor at the last successful chunk so the next cycle
			// retries exactly this range; chunks scanned before the failure
			// stay counted.
			r.logger.Warn().
				Err(err).
				Uint64("from_block", from).
				Uint64("to_block", to).
				Msg("Failed to fetch state sync logs")
			return
		}

		r.recordStateSyncLogs(receiver, logs)
		r.stateSyncCursor = to
	}
}

// recordStateSyncLogs decodes, validates and deduplicates raw logs, adding
// the well-formed events to this cycle's accumulation. Malformed, removed or
// wrongly-addressed logs are logged and skipped: they are never counted.
func (r *RPCProvider) recordStateSyncLogs(receiver common.Address, logs []types.Log) {
	for i := range logs {
		l := &logs[i]

		// The log filter already constrains the emitter, but do not blindly
		// trust the RPC to have honored it.
		if l.Address != receiver {
			r.logMalformedStateSyncLog(l, "unexpected emitter address")
			continue
		}

		// A removed log is a reorg artifact, not a chain event.
		if l.Removed {
			continue
		}

		key := stateSyncLogKey{blockHash: l.BlockHash, logIndex: uint32(l.Index)}
		if r.stateSyncSeen.seen(key) {
			continue
		}

		if len(l.Topics) == 0 {
			r.logMalformedStateSyncLog(l, "missing event signature topic")
			continue
		}

		switch l.Topics[0] {
		case stateCommittedTopic:
			commit, ok := decodeStateSyncCommit(l)
			if !ok {
				r.logMalformedStateSyncLog(l, "invalid StateCommitted encoding")
				continue
			}
			r.stateSyncEvents.Commits = append(r.stateSyncEvents.Commits, commit)
			r.stateSyncSeen.add(key)

		case stateSyncReplayTopic:
			replay, ok := decodeStateSyncReplay(l)
			if !ok {
				r.logMalformedStateSyncLog(l, "invalid StateSyncReplay encoding")
				continue
			}
			r.stateSyncEvents.Replays = append(r.stateSyncEvents.Replays, replay)
			r.stateSyncSeen.add(key)

		default:
			// The log filter asks for exactly our two topics; anything else
			// here would mean the RPC returned unrequested logs.
			r.logMalformedStateSyncLog(l, "unexpected event signature")
		}
	}
}

func (r *RPCProvider) logMalformedStateSyncLog(l *types.Log, reason string) {
	r.logger.Warn().
		Str("reason", reason).
		Uint64("block_number", l.BlockNumber).
		Str("block_hash", l.BlockHash.Hex()).
		Uint32("log_index", uint32(l.Index)).
		Msg("Discarding malformed state sync log")
}

// decodeStateSyncCommit decodes and validates a
// StateCommitted(uint256 indexed stateId, bool success) log: exactly two
// topics (signature and indexed stateId) and one 32-byte data word holding a
// canonical ABI boolean. State IDs start at 1 and must fit in uint64.
func decodeStateSyncCommit(l *types.Log) (*observer.StateSyncCommit, bool) {
	if len(l.Topics) != 2 || len(l.Data) != 32 {
		return nil, false
	}

	stateId, ok := decodeStateSyncId(l.Topics[1])
	if !ok {
		return nil, false
	}

	success, ok := decodeStateSyncBool(l.Data)
	if !ok {
		return nil, false
	}

	return &observer.StateSyncCommit{
		StateId:     stateId,
		Success:     success,
		BlockNumber: l.BlockNumber,
		BlockHash:   l.BlockHash,
		LogIndex:    uint32(l.Index),
	}, true
}

// decodeStateSyncReplay decodes and validates a
// StateSyncReplay(uint256 indexed stateId) log: exactly two topics (signature
// and indexed stateId) and no data.
func decodeStateSyncReplay(l *types.Log) (*observer.StateSyncReplay, bool) {
	if len(l.Topics) != 2 || len(l.Data) != 0 {
		return nil, false
	}

	stateId, ok := decodeStateSyncId(l.Topics[1])
	if !ok {
		return nil, false
	}

	return &observer.StateSyncReplay{
		StateId:     stateId,
		BlockNumber: l.BlockNumber,
		BlockHash:   l.BlockHash,
		LogIndex:    uint32(l.Index),
	}, true
}

// decodeStateSyncId validates the indexed stateId topic: state IDs start at
// 1 and must fit in uint64.
func decodeStateSyncId(topic common.Hash) (uint64, bool) {
	id := new(big.Int).SetBytes(topic[:])
	if id.Sign() <= 0 || !id.IsUint64() {
		return 0, false
	}
	return id.Uint64(), true
}

// decodeStateSyncBool accepts only the canonical ABI encodings of false and
// true: a 32-byte word that is all zeroes, or all zeroes with a single 1 in
// the last byte.
func decodeStateSyncBool(word []byte) (bool, bool) {
	for _, b := range word[:31] {
		if b != 0 {
			return false, false
		}
	}
	switch word[31] {
	case 0:
		return false, true
	case 1:
		return true, true
	default:
		return false, false
	}
}
