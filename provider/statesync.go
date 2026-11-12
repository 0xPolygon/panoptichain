package provider

import (
	"context"
	"math/big"
	"time"

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

	// stateSyncMaxChunksPerCycle bounds how many chunks one refresh cycle
	// will scan. Each call is already bounded by stateSyncScanChunk, but
	// without a per-cycle cap a large catch-up backlog (e.g. after an
	// outage) could consume the whole cycle. The scan yields at the last
	// completed chunk and the next cycle resumes from there, so the other
	// collectors sharing the cycle keep their headroom. 50 chunks = 5000
	// blocks ≈ 2.8 hours of Bor blocks per cycle.
	stateSyncMaxChunksPerCycle = 50

	// stateSyncScanTimeout bounds the whole scan's share of a refresh cycle,
	// following the account-balance sweep convention: a subordinate deadline
	// so a slow log endpoint cannot starve the checkpoint, validator and
	// account collectors that share the cycle's context.
	stateSyncScanTimeout = 15 * time.Second

	// stateSyncSeenCap bounds the deduplication set. It comfortably covers
	// many cycles of ordinary state-sync traffic.
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
// delivered log comes back in an overlapping range, a duplicate log entry, or
// a retried RPC response.
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
// StateSyncReplay logs when both the per-RPC scan flag and a state sync
// receiver are configured. The observer's enablement controls metric
// consumption independently; it does not enable this RPC work. This is a
// metrics-only first slice: all scan state is process-local. The scan runs
// under a subordinate timeout so a large backlog or a slow log endpoint
// cannot starve the collectors that share this refresh cycle.
func (r *RPCProvider) refreshStateSyncEvents(ctx context.Context, c *ethclient.Client) {
	if !r.stateSyncEventsEnabled || r.contracts.StateSyncReceiverAddress == nil {
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

	scanCtx, cancel := context.WithTimeout(ctx, stateSyncScanTimeout)
	defer cancel()

	r.scanStateSyncLogs(scanCtx, receiver, r.finalizedHeight, fetch)
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
//   - Response validation: every log in a fetched response must match what
//     was requested — the StateReceiver emitter, the exact numeric range, a
//     non-zero block hash, no removed flag and a decodable event. Any
//     inconsistent log rejects the whole chunk: it is neither counted nor
//     skipped, and its range is retried next cycle. A chunk is committed
//     (events + dedup keys + cursor) only when it validates in full, so a
//     retry can never double count a partially valid response.
//   - Cursor: the cursor only advances after a fully successful chunk scan.
//     A fetch error or a rejected chunk aborts the scan; the next cycle
//     retries the same range, so blocks are never skipped and successfully
//     scanned chunks are never re-scanned.
//   - Budget: the scan performs at most stateSyncMaxChunksPerCycle chunk
//     fetches per cycle and yields immediately on context cancellation
//     (including its subordinate scan deadline). Catch-up resumes at the
//     last completed chunk next cycle, leaving headroom for the other
//     collectors in the cycle.
//   - Bootstrap: the first scan records the current finalized head as its
//     cursor without fetching or counting historical events. Only blocks
//     finalized after bootstrap are counted, avoiding a restart burst of old
//     outcomes. While this process stays running, the same cursor retains and
//     catches up on any backlog; a later restart establishes a new head.
//   - Deduplication: a bounded (block hash, log index) set prevents the same
//     log from being counted twice across overlapping ranges or retried
//     responses.
//   - Restart: the cursor, dedup set and counters are process-local. A restart
//     bootstraps at the new finalized head and Prometheus counters reset with
//     the process; no historical backfill or durable exactly-once delivery is
//     claimed.
//
// getFilterOpts is not reusable here: its latest-head/previous-height bounds
// move as a pair, while this cursor needs independent finalized upper bounds
// and fixed chunks that commit one at a time.
func (r *RPCProvider) scanStateSyncLogs(ctx context.Context, receiver common.Address, finalized uint64, fetch stateSyncLogFetcher) {
	// Each cycle publishes only what this scan accumulates.
	r.stateSyncEvents = &observer.StateSyncEvents{}

	// The finalized head was not fetched this cycle; retry next cycle.
	if finalized == 0 {
		return
	}

	if !r.stateSyncBootstrapped {
		// Counters start at zero with the process. Do not add a lookback's
		// historical events to them as a burst at startup; begin with blocks
		// finalized after this head instead. The cursor still catches up on
		// gaps that occur while this process remains running.
		r.stateSyncCursor = finalized
		r.stateSyncBootstrapped = true

		r.logger.Info().
			Uint64("cursor_block", finalized).
			Msg("Bootstrapping state sync event scan after finalized head")
		return
	}

	// A lagging RPC view can report a finalized head below our cursor. Wait
	// for it to catch up rather than rescanning, so counting stays monotonic
	// without skipping blocks.
	if finalized <= r.stateSyncCursor {
		return
	}

	chunks := 0
	for from := r.stateSyncCursor + 1; from <= finalized; from += stateSyncScanChunk {
		// Yield on cancellation (including the subordinate scan deadline).
		// The cursor already records the last completed chunk, so the next
		// cycle resumes right here without skipping or recounting.
		if ctx.Err() != nil {
			return
		}

		// Bound this scan's share of the refresh cycle; catch-up resumes
		// from the last completed chunk next cycle.
		if chunks >= stateSyncMaxChunksPerCycle {
			r.logger.Info().
				Uint64("resume_block", from).
				Int("chunks", chunks).
				Msg("State sync scan reached its per-cycle chunk budget; continuing next cycle")
			return
		}

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

		staged, ok := r.decodeStateSyncChunk(receiver, from, to, logs)
		if !ok {
			// The response was inconsistent with the request. Like a fetch
			// failure, this leaves the cursor at the last successful chunk so
			// the range is retried — never silently skipped, and nothing from
			// this chunk was counted or recorded in the dedup set.
			r.logger.Warn().
				Uint64("from_block", from).
				Uint64("to_block", to).
				Msg("Inconsistent state sync log response; will retry this range next cycle")
			return
		}

		r.commitStateSyncChunk(staged)
		r.stateSyncCursor = to
		chunks++
	}
}

// decodeStateSyncChunk validates an entire fetched response against what was
// requested and stages the well-formed events for commit. It returns ok=false
// when any returned log is inconsistent — a wrong emitter, a removed flag, a
// block number outside [from, to], a zero block hash, an unknown signature
// or an undecodable event — in which case the caller must reject the whole
// chunk and retry its range. Every returned log is decoded and validated
// before deduplication can skip it, so a malformed log can never hide behind
// a duplicate key: only a well-formed duplicate is suppressed as already
// counted (an overlapping range or a re-delivered response).
func (r *RPCProvider) decodeStateSyncChunk(receiver common.Address, from, to uint64, logs []types.Log) (*observer.StateSyncEvents, bool) {
	staged := &observer.StateSyncEvents{}
	chunkSeen := make(map[stateSyncLogKey]struct{})

	for i := range logs {
		l := &logs[i]

		// The log filter already constrains the emitter and topics, but do
		// not blindly trust the RPC to have honored them — or the range.
		if l.Address != receiver {
			r.logMalformedStateSyncLog(l, "unexpected emitter address")
			return nil, false
		}

		// A removed log is a reorg artifact; a finalized range has none.
		if l.Removed {
			r.logMalformedStateSyncLog(l, "removed log in a finalized range")
			return nil, false
		}

		// The response must stay inside the requested numeric range.
		if l.BlockNumber < from || l.BlockNumber > to {
			r.logMalformedStateSyncLog(l, "log outside the requested range")
			return nil, false
		}

		// A mined log must carry the hash of its block.
		if l.BlockHash == (common.Hash{}) {
			r.logMalformedStateSyncLog(l, "missing block hash")
			return nil, false
		}

		if len(l.Topics) == 0 {
			r.logMalformedStateSyncLog(l, "missing event signature topic")
			return nil, false
		}

		// Decode every returned event before deduplication gets a chance
		// to skip it: a malformed log that reuses a counted or in-response
		// key must reject the chunk like any other inconsistent log, not
		// slip past validation as a "duplicate".
		var (
			commit *observer.StateSyncCommit
			replay *observer.StateSyncReplay
		)
		switch l.Topics[0] {
		case stateCommittedTopic:
			c, ok := decodeStateSyncCommit(l)
			if !ok {
				r.logMalformedStateSyncLog(l, "invalid StateCommitted encoding")
				return nil, false
			}
			commit = c

		case stateSyncReplayTopic:
			rp, ok := decodeStateSyncReplay(l)
			if !ok {
				r.logMalformedStateSyncLog(l, "invalid StateSyncReplay encoding")
				return nil, false
			}
			replay = rp

		default:
			// The log filter asks for exactly our two topics; anything else
			// here would mean the RPC returned unrequested logs.
			r.logMalformedStateSyncLog(l, "unexpected event signature")
			return nil, false
		}

		key := stateSyncLogKey{blockHash: l.BlockHash, logIndex: uint32(l.Index)}
		if _, dup := chunkSeen[key]; dup || r.stateSyncSeen.seen(key) {
			// A well-formed duplicate: already counted by an earlier chunk
			// or delivered twice within this response. Overlap is not an
			// inconsistency.
			continue
		}

		if commit != nil {
			staged.Commits = append(staged.Commits, commit)
		} else {
			staged.Replays = append(staged.Replays, replay)
		}
		chunkSeen[key] = struct{}{}
	}

	return staged, true
}

// commitStateSyncChunk moves a fully validated chunk's events into this
// cycle's accumulation and records their log keys in the dedup set. It runs
// only after decodeStateSyncChunk accepted the chunk in full, so a retried
// range can never double count a partially valid response.
func (r *RPCProvider) commitStateSyncChunk(staged *observer.StateSyncEvents) {
	for _, commit := range staged.Commits {
		r.stateSyncEvents.Commits = append(r.stateSyncEvents.Commits, commit)
		r.stateSyncSeen.add(stateSyncLogKey{blockHash: commit.BlockHash, logIndex: commit.LogIndex})
	}
	for _, replay := range staged.Replays {
		r.stateSyncEvents.Replays = append(r.stateSyncEvents.Replays, replay)
		r.stateSyncSeen.add(stateSyncLogKey{blockHash: replay.BlockHash, logIndex: replay.LogIndex})
	}
}

// logMalformedStateSyncLog records why a fetched response is inconsistent.
// The offending log's whole chunk is rejected and its range retried next
// cycle; the log itself is never counted.
func (r *RPCProvider) logMalformedStateSyncLog(l *types.Log, reason string) {
	r.logger.Warn().
		Str("reason", reason).
		Uint64("block_number", l.BlockNumber).
		Str("block_hash", l.BlockHash.Hex()).
		Uint32("log_index", uint32(l.Index)).
		Msg("Rejecting inconsistent state sync log")
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
