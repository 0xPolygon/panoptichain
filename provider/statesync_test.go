package provider

import (
	"context"
	"errors"
	"math"
	"math/big"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/0xPolygon/panoptichain/observer"
)

// The configured StateReceiver (0x…1001 on Polygon mainnet and Amoy) and a
// stand-in for any other emitter.
var (
	stateReceiverAddress = common.HexToAddress("0x0000000000000000000000000000000000001001")
	otherEmitterAddress  = common.HexToAddress("0x000000000000000000000000000000000000dead")
)

// syntheticTxHash stands in for Bor's synthetic state-sync transaction. The
// same hash carrying several commit logs mirrors the multi-ID synthetic
// receipts observed on mainnet. The receipt's status (always 1, gasUsed 0) is
// deliberately not modeled: the collector counts the commit event's boolean,
// not receipt status, so a failed receiver call inside a successful synthetic
// receipt still counts as outcome=failure.
var syntheticTxHash = common.HexToHash("0xe849ab3cba807c960d7219e6b189ce344b522e215b2577877b67540e1b2b14a0")

func stateSyncBlockHash(block uint64) common.Hash {
	return common.BigToHash(new(big.Int).SetUint64(block))
}

func stateIdTopic(stateId uint64) common.Hash {
	return common.BigToHash(new(big.Int).SetUint64(stateId))
}

func commitLog(block, stateId uint64, index uint, success bool) types.Log {
	data := make([]byte, 32)
	if success {
		data[31] = 1
	}

	return types.Log{
		Address:     stateReceiverAddress,
		Topics:      []common.Hash{stateCommittedTopic, stateIdTopic(stateId)},
		Data:        data,
		BlockNumber: block,
		BlockHash:   stateSyncBlockHash(block),
		TxHash:      syntheticTxHash,
		Index:       index,
	}
}

func replayLog(block, stateId uint64, index uint) types.Log {
	return types.Log{
		Address:     stateReceiverAddress,
		Topics:      []common.Hash{stateSyncReplayTopic, stateIdTopic(stateId)},
		BlockNumber: block,
		BlockHash:   stateSyncBlockHash(block),
		TxHash:      syntheticTxHash,
		Index:       index,
	}
}

func newStateSyncProvider(lookback uint64) *RPCProvider {
	return &RPCProvider{
		logger:        NewLogger(nil, "test"),
		blockLookBack: lookback,
		stateSyncSeen: newStateSyncSeenLogs(),
		// recordStateSyncLogs is tested directly, without a full scan cycle
		// to initialize the accumulator.
		stateSyncEvents: &observer.StateSyncEvents{},
	}
}

// rangeRecorder captures the ranges a scan requests and answers them from a
// per-range log map. Ranges without an entry return no logs, which is what an
// empty (zero-traffic) scan looks like.
type rangeRecorder struct {
	ranges [][2]uint64
	logs   map[[2]uint64][]types.Log
}

func (rec *rangeRecorder) fetch(_ context.Context, from, to uint64) ([]types.Log, error) {
	rec.ranges = append(rec.ranges, [2]uint64{from, to})
	return rec.logs[[2]uint64{from, to}], nil
}

func commitIds(events *observer.StateSyncEvents) []uint64 {
	ids := make([]uint64, 0, len(events.Commits))
	for _, c := range events.Commits {
		ids = append(ids, c.StateId)
	}
	return ids
}

func replayIds(events *observer.StateSyncEvents) []uint64 {
	ids := make([]uint64, 0, len(events.Replays))
	for _, r := range events.Replays {
		ids = append(ids, r.StateId)
	}
	return ids
}

func TestDecodeStateSyncCommit(t *testing.T) {
	tests := []struct {
		name    string
		log     types.Log
		wantId  uint64
		wantOK  bool
		success bool
	}{
		{name: "success commit decodes", log: commitLog(100, 42, 0, true), wantId: 42, wantOK: true, success: true},
		{name: "failure commit decodes", log: commitLog(100, 42, 0, false), wantId: 42, wantOK: true},
		{name: "max uint64 state id decodes", log: commitLog(100, math.MaxUint64, 0, true), wantId: math.MaxUint64, wantOK: true, success: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commit, ok := decodeStateSyncCommit(&tt.log)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if !ok {
				return
			}
			if commit.StateId != tt.wantId || commit.Success != tt.success {
				t.Errorf("commit = (id %d, success %v), want (id %d, success %v)",
					commit.StateId, commit.Success, tt.wantId, tt.success)
			}
		})
	}
}

func TestDecodeStateSyncCommit_Malformed(t *testing.T) {
	tests := []struct {
		name string
		log  types.Log
	}{
		{
			name: "missing indexed stateId topic",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateCommittedTopic},
				Data:    make([]byte, 32),
			},
		},
		{
			name: "extra indexed topic",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateCommittedTopic, stateIdTopic(1), stateIdTopic(2)},
				Data:    make([]byte, 32),
			},
		},
		{
			name: "short data word",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateCommittedTopic, stateIdTopic(1)},
				Data:    make([]byte, 31),
			},
		},
		{
			name: "long data word",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateCommittedTopic, stateIdTopic(1)},
				Data:    make([]byte, 33),
			},
		},
		{
			name: "non-canonical boolean value two",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateCommittedTopic, stateIdTopic(1)},
				Data:    append(make([]byte, 31), 2),
			},
		},
		{
			name: "non-canonical boolean with dirty high bytes",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateCommittedTopic, stateIdTopic(1)},
				Data:    append([]byte{1}, make([]byte, 31)...),
			},
		},
		{
			name: "zero state id",
			log:  commitLog(100, 0, 0, true),
		},
		{
			name: "state id above max uint64",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics: []common.Hash{
					stateCommittedTopic,
					common.HexToHash("0x0000000000000000000000000000000100000000000000000000000000000000"),
				},
				Data: make([]byte, 32),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if commit, ok := decodeStateSyncCommit(&tt.log); ok {
				t.Fatalf("expected malformed log to be rejected, got commit %+v", commit)
			}
		})
	}
}

func TestDecodeStateSyncReplay(t *testing.T) {
	tests := []struct {
		name   string
		log    types.Log
		wantId uint64
		wantOK bool
	}{
		{name: "replay decodes", log: replayLog(100, 42, 0), wantId: 42, wantOK: true},
		{
			name: "replay with data is rejected",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateSyncReplayTopic, stateIdTopic(42)},
				Data:    make([]byte, 32),
			},
		},
		{
			name: "replay missing indexed topic",
			log: types.Log{
				Address: stateReceiverAddress,
				Topics:  []common.Hash{stateSyncReplayTopic},
			},
		},
		{name: "replay with zero state id", log: replayLog(100, 0, 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replay, ok := decodeStateSyncReplay(&tt.log)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if !ok {
				return
			}
			if replay.StateId != tt.wantId {
				t.Errorf("replay state id = %d, want %d", replay.StateId, tt.wantId)
			}
		})
	}
}

func TestRecordStateSyncLogs(t *testing.T) {
	r := newStateSyncProvider(1000)

	logs := []types.Log{
		commitLog(100, 1, 0, true),
		// A failed receiver call inside a successful synthetic receipt: the
		// receipt status is not part of the log, only the boolean is.
		commitLog(100, 2, 1, false),
		replayLog(101, 2, 0),
		// Several state IDs delivered in one synthetic transaction.
		commitLog(102, 3, 0, true),
		commitLog(102, 4, 1, true),
		commitLog(102, 5, 2, false),
		// Never counted: reorg artifact, wrong emitter, malformed.
		func() types.Log {
			l := commitLog(103, 6, 0, true)
			l.Removed = true
			return l
		}(),
		func() types.Log {
			l := commitLog(103, 7, 1, true)
			l.Address = otherEmitterAddress
			return l
		}(),
		func() types.Log {
			l := commitLog(103, 8, 2, true)
			l.Data = append(make([]byte, 31), 2)
			return l
		}(),
		// No topics at all.
		{Address: stateReceiverAddress, BlockNumber: 103, BlockHash: stateSyncBlockHash(103), Index: 3},
	}

	r.recordStateSyncLogs(stateReceiverAddress, logs)

	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1, 2, 3, 4, 5}) {
		t.Errorf("commit ids = %v, want [1 2 3 4 5]", got)
	}
	if got := replayIds(r.stateSyncEvents); !slices.Equal(got, []uint64{2}) {
		t.Errorf("replay ids = %v, want [2]", got)
	}
	for _, c := range r.stateSyncEvents.Commits {
		if c.BlockHash != stateSyncBlockHash(c.BlockNumber) {
			t.Errorf("commit %d block hash not carried", c.StateId)
		}
	}

	// Re-delivering the same logs (an overlapping range, a retried response or
	// a republished message) must not count anything twice.
	r.stateSyncEvents = &observer.StateSyncEvents{}
	r.recordStateSyncLogs(stateReceiverAddress, logs)
	if got := commitIds(r.stateSyncEvents); len(got) != 0 {
		t.Errorf("duplicate delivery counted %d commits, want 0", len(got))
	}
	if got := replayIds(r.stateSyncEvents); len(got) != 0 {
		t.Errorf("duplicate delivery counted %d replays, want 0", len(got))
	}
}

func TestRecordStateSyncLogs_DuplicateWithinOneResponse(t *testing.T) {
	r := newStateSyncProvider(1000)
	dup := commitLog(100, 9, 3, true)

	r.recordStateSyncLogs(stateReceiverAddress, []types.Log{dup, dup})

	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{9}) {
		t.Errorf("commit ids = %v, want [9]", got)
	}
}

// A fresh provider bootstraps a fixed lookback below the current finalized
// head and scans forward in fixed chunks; an empty scan still advances the
// cursor (a zero-traffic scan is a successful scan).
func TestScanStateSyncLogs_BootstrapAndEmptyScan(t *testing.T) {
	r := newStateSyncProvider(1000)
	rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{}}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec.fetch)

	wantRanges := [][2]uint64{{1, 100}, {101, 200}, {201, 250}}
	if !slices.Equal(rec.ranges, wantRanges) {
		t.Errorf("ranges = %v, want %v", rec.ranges, wantRanges)
	}
	if r.stateSyncCursor != 250 {
		t.Errorf("cursor = %d, want 250", r.stateSyncCursor)
	}
	if len(r.stateSyncEvents.Commits) != 0 || len(r.stateSyncEvents.Replays) != 0 {
		t.Errorf("empty scan produced events")
	}

	// A second cycle with the same finalized head does nothing.
	rec.ranges = nil
	r.stateSyncEvents = &observer.StateSyncEvents{}
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec.fetch)
	if len(rec.ranges) != 0 {
		t.Errorf("caught-up scan refetched %v", rec.ranges)
	}
}

func TestScanStateSyncLogs_BootstrapLookback(t *testing.T) {
	r := newStateSyncProvider(1000)
	rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{}}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 2500, rec.fetch)

	if len(rec.ranges) == 0 {
		t.Fatal("no ranges scanned")
	}
	if got := rec.ranges[0]; got != [2]uint64{1500, 1599} {
		t.Errorf("first range = %v, want [1500 1599]", got)
	}
	if r.stateSyncCursor != 2500 {
		t.Errorf("cursor = %d, want 2500", r.stateSyncCursor)
	}
}

// A chunk failure must not skip blocks: the cursor stays at the last
// successful chunk, and the next cycle resumes the failed range.
func TestScanStateSyncLogs_PartialFailureRetriesWithoutSkipping(t *testing.T) {
	r := newStateSyncProvider(250)

	chunkOne := [2]uint64{300, 399}
	chunkTwo := [2]uint64{400, 499}
	chunkThree := [2]uint64{500, 550}
	logs := map[[2]uint64][]types.Log{
		chunkOne:   {commitLog(310, 1, 0, true), commitLog(350, 2, 0, false)},
		chunkTwo:   {commitLog(420, 3, 0, false)},
		chunkThree: {replayLog(510, 2, 0)},
	}

	var ranges [][2]uint64
	fail := true
	fetch := func(_ context.Context, from, to uint64) ([]types.Log, error) {
		key := [2]uint64{from, to}
		ranges = append(ranges, key)
		if fail && key == chunkTwo {
			return nil, errors.New("rpc is down")
		}
		return logs[key], nil
	}

	// Cycle 1: chunk two fails after chunk one succeeded.
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 550, fetch)
	if r.stateSyncCursor != 399 {
		t.Fatalf("cursor after partial failure = %d, want 399", r.stateSyncCursor)
	}
	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1, 2}) {
		t.Errorf("cycle 1 commits = %v, want [1 2] (successful chunks still count)", got)
	}

	// Cycle 2: the scan resumes exactly at the failed chunk and completes.
	fail = false
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 550, fetch)
	chunkOneFetches := 0
	for _, rg := range ranges {
		if rg == chunkOne {
			chunkOneFetches++
		}
	}
	if chunkOneFetches != 1 {
		t.Errorf("chunk one fetched %d times, want 1 (no re-scan of successful chunks)", chunkOneFetches)
	}
	if got := ranges[len(ranges)-2:]; !slices.Equal(got, [][2]uint64{chunkTwo, chunkThree}) {
		t.Errorf("cycle 2 ranges = %v, want %v", got, [][2]uint64{chunkTwo, chunkThree})
	}
	if r.stateSyncCursor != 550 {
		t.Errorf("cursor after retry = %d, want 550", r.stateSyncCursor)
	}
	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{3}) {
		t.Errorf("cycle 2 commits = %v, want [3] (no re-count of chunk one)", got)
	}
	if got := replayIds(r.stateSyncEvents); !slices.Equal(got, []uint64{2}) {
		t.Errorf("cycle 2 replays = %v, want [2]", got)
	}
}

// An RPC that returns the same logs for overlapping ranges (or re-delivers
// them across retries) must not cause double counting: the (block hash, log
// index) dedup set counts each log once.
func TestScanStateSyncLogs_OverlappingRangesDeduplicate(t *testing.T) {
	r := newStateSyncProvider(100)

	allLogs := []types.Log{
		commitLog(50, 1, 0, true),
		commitLog(150, 2, 0, false),
		replayLog(180, 1, 0),
	}

	// A misbehaving fetcher: every range returns every log.
	fetch := func(_ context.Context, from, to uint64) ([]types.Log, error) {
		return allLogs, nil
	}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 200, fetch)

	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1, 2}) {
		t.Errorf("commit ids = %v, want [1 2] (overlaps deduplicated)", got)
	}
	if got := replayIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1}) {
		t.Errorf("replay ids = %v, want [1]", got)
	}
}

// A lagging RPC view can report a finalized head below the cursor. The scan
// waits rather than rescanning, so counting stays monotonic without skipping.
// Near-head reorgs therefore cannot double count: only finalized heights are
// ever scanned. (A reorg of already-finalized blocks is not detected by this
// collector; it is a chain integrity failure.)
func TestScanStateSyncLogs_FinalizedBehindCursorStalls(t *testing.T) {
	r := newStateSyncProvider(100)
	rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{}}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 500, rec.fetch)

	rec.ranges = nil
	r.stateSyncEvents = &observer.StateSyncEvents{}
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 400, rec.fetch)

	if len(rec.ranges) != 0 {
		t.Errorf("stalled scan refetched %v", rec.ranges)
	}
	if r.stateSyncCursor != 500 {
		t.Errorf("cursor moved backwards: %d", r.stateSyncCursor)
	}
}

// All scan state is process-local: a restart builds a fresh provider that
// re-bootstraps the same lookback window (counters reset with the process
// too, so this is not double counting in one series — it is the documented
// restart behavior).
func TestScanStateSyncLogs_RestartRebootstrap(t *testing.T) {
	first := newStateSyncProvider(1000)
	rec1 := &rangeRecorder{logs: map[[2]uint64][]types.Log{}}
	first.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec1.fetch)

	second := newStateSyncProvider(1000)
	rec2 := &rangeRecorder{logs: map[[2]uint64][]types.Log{}}
	second.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec2.fetch)

	if !slices.Equal(rec1.ranges, rec2.ranges) {
		t.Errorf("restart ranges %v differ from first bootstrap %v", rec2.ranges, rec1.ranges)
	}
}

// A finalized head of zero means the finalized header was not fetched this
// cycle; the scan must do nothing and try again next cycle.
func TestScanStateSyncLogs_NoFinalizedHead(t *testing.T) {
	r := newStateSyncProvider(1000)
	rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{}}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 0, rec.fetch)

	if len(rec.ranges) != 0 {
		t.Errorf("scanned without a finalized head: %v", rec.ranges)
	}
	if r.stateSyncBootstrapped {
		t.Error("scan bootstrapped without a finalized head")
	}
}

func TestStateSyncSeenLogs_Bound(t *testing.T) {
	s := newStateSyncSeenLogs()

	for i := uint64(0); i < stateSyncSeenCap*2; i++ {
		s.add(stateSyncLogKey{blockHash: stateSyncBlockHash(i), logIndex: 0})
	}

	if len(s.keys) > stateSyncSeenCap {
		t.Errorf("dedup set size = %d, want at most %d", len(s.keys), stateSyncSeenCap)
	}
	// The most recent entries must still be seen.
	if !s.seen(stateSyncLogKey{blockHash: stateSyncBlockHash(stateSyncSeenCap*2 - 1), logIndex: 0}) {
		t.Error("latest key was evicted")
	}
	// The oldest entries must have been evicted.
	if s.seen(stateSyncLogKey{blockHash: stateSyncBlockHash(0), logIndex: 0}) {
		t.Error("oldest key was retained past the cap")
	}
}
