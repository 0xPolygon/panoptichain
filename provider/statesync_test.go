package provider

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
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

// The full decode/validate/dedup path over one fetched response: true and
// false commits, a failed receiver inside a successful synthetic receipt,
// replays, several state IDs delivered in one synthetic transaction, and a
// log delivered twice within one response.
func TestDecodeStateSyncChunk(t *testing.T) {
	r := newStateSyncProvider(1000)

	logs := []types.Log{
		commitLog(100, 1, 0, true),
		// A failed receiver call inside a successful synthetic receipt: the
		// receipt status is not part of the log, only the boolean is.
		commitLog(100, 2, 1, false),
		replayLog(101, 2, 0),
		// Several state IDs delivered in one synthetic transaction, as
		// distinct logs (distinct log indices).
		commitLog(102, 3, 0, true),
		commitLog(102, 4, 1, true),
		commitLog(102, 5, 2, false),
		// The same log delivered twice within one response.
		commitLog(102, 5, 2, false),
	}

	staged, ok := r.decodeStateSyncChunk(stateReceiverAddress, 100, 102, logs)
	if !ok {
		t.Fatal("well-formed response rejected")
	}
	if got := commitIds(staged); !slices.Equal(got, []uint64{1, 2, 3, 4, 5}) {
		t.Errorf("commit ids = %v, want [1 2 3 4 5]", got)
	}
	if got := replayIds(staged); !slices.Equal(got, []uint64{2}) {
		t.Errorf("replay ids = %v, want [2]", got)
	}
	for _, c := range staged.Commits {
		if c.BlockHash != stateSyncBlockHash(c.BlockNumber) {
			t.Errorf("commit %d block hash not carried", c.StateId)
		}
	}

	// After commit, re-delivering the same logs (an overlapping range or a
	// retried response) is valid but stages nothing: every key is already in
	// the dedup set, so nothing is counted twice.
	r.commitStateSyncChunk(staged)
	staged2, ok := r.decodeStateSyncChunk(stateReceiverAddress, 100, 102, logs)
	if !ok {
		t.Fatal("re-delivered response rejected")
	}
	if len(staged2.Commits) != 0 || len(staged2.Replays) != 0 {
		t.Errorf("duplicate delivery staged %d commits and %d replays, want 0",
			len(staged2.Commits), len(staged2.Replays))
	}
}

// A StateSyncReplay in an ordinary transaction (a real transaction hash, not
// Bor's synthetic state-sync receipt) is the same event and decodes the same.
func TestDecodeStateSyncChunk_OrdinaryTransactionReplay(t *testing.T) {
	r := newStateSyncProvider(1000)

	l := replayLog(100, 42, 0)
	l.TxHash = common.HexToHash("0x00000000000000000000000000000000000000000000000000000000feedbeef")

	staged, ok := r.decodeStateSyncChunk(stateReceiverAddress, 100, 100, []types.Log{l})
	if !ok {
		t.Fatal("ordinary-transaction replay rejected")
	}
	if got := replayIds(staged); !slices.Equal(got, []uint64{42}) {
		t.Errorf("replay ids = %v, want [42]", got)
	}
}

// Any single inconsistent log rejects the whole chunk — it is never counted
// and never silently skipped. A well-formed log in the same response must not
// be partially committed either: the next decode of that log alone still
// stages it, proving no dedup or event state leaked from the rejected chunk.
func TestDecodeStateSyncChunk_InconsistentResponseRejected(t *testing.T) {
	inRange := commitLog(100, 1, 0, true)

	tests := []struct {
		name string
		log  types.Log
	}{
		// The bad logs carry a distinct log index from the in-range commit so
		// they exercise the rejection paths, not the dedup skip.
		{name: "log above the requested range", log: commitLog(101, 6, 1, true)},
		{name: "log below the requested range", log: commitLog(99, 6, 1, true)},
		{
			name: "removed log in a finalized range",
			log: func() types.Log {
				l := commitLog(100, 6, 1, true)
				l.Removed = true
				return l
			}(),
		},
		{
			name: "zero block hash",
			log: func() types.Log {
				l := commitLog(100, 6, 1, true)
				l.BlockHash = common.Hash{}
				return l
			}(),
		},
		{
			name: "wrong emitter address",
			log: func() types.Log {
				l := commitLog(100, 6, 1, true)
				l.Address = otherEmitterAddress
				return l
			}(),
		},
		// No topics at all.
		{
			name: "no topics",
			log: types.Log{
				Address:     stateReceiverAddress,
				BlockNumber: 100,
				BlockHash:   stateSyncBlockHash(100),
				Index:       3,
			},
		},
		{
			name: "unexpected event signature",
			log: types.Log{
				Address:     stateReceiverAddress,
				Topics:      []common.Hash{stateIdTopic(6)},
				BlockNumber: 100,
				BlockHash:   stateSyncBlockHash(100),
				Index:       1,
			},
		},
		{
			name: "invalid commit boolean",
			log: func() types.Log {
				l := commitLog(100, 6, 1, true)
				l.Data = append(make([]byte, 31), 2)
				return l
			}(),
		},
		// A malformed log reusing the in-range log's (block hash, log
		// index) key: deduplication must not let it bypass validation.
		{
			name: "malformed duplicate of the in-range log",
			log: func() types.Log {
				l := commitLog(100, 6, 0, true)
				l.Data = append(make([]byte, 31), 2)
				return l
			}(),
		},
		{
			name: "replay with data",
			log: func() types.Log {
				l := replayLog(100, 6, 1)
				l.Data = make([]byte, 32)
				return l
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newStateSyncProvider(1000)

			staged, ok := r.decodeStateSyncChunk(stateReceiverAddress, 100, 100, []types.Log{inRange, tt.log})
			if ok {
				t.Fatal("inconsistent response accepted")
			}
			if staged != nil && (len(staged.Commits) > 0 || len(staged.Replays) > 0) {
				t.Fatal("rejected chunk staged events")
			}

			// Nothing from the rejected chunk may have been committed: the
			// well-formed log is staged again by a clean response.
			staged, ok = r.decodeStateSyncChunk(stateReceiverAddress, 100, 100, []types.Log{inRange})
			if !ok {
				t.Fatal("clean response rejected after an inconsistent one")
			}
			if got := commitIds(staged); !slices.Equal(got, []uint64{1}) {
				t.Errorf("commit ids after retry = %v, want [1] (no partial commit leaked)", got)
			}
		})
	}
}

// A fresh provider records the current finalized head as its cursor without
// fetching or counting historical logs. Only blocks finalized after bootstrap
// are scanned and counted.
func TestScanStateSyncLogs_BootstrapSkipsHistoricalAndCountsNextFinalized(t *testing.T) {
	r := newStateSyncProvider(1000)
	rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{
		{250, 250}: {commitLog(250, 1, 0, false)},
		{251, 251}: {commitLog(251, 2, 0, true)},
	}}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec.fetch)

	if len(rec.ranges) != 0 {
		t.Fatalf("fresh boot fetched historical ranges %v, want none", rec.ranges)
	}
	if r.stateSyncCursor != 250 {
		t.Errorf("cursor = %d, want finalized head 250", r.stateSyncCursor)
	}
	if len(r.stateSyncEvents.Commits) != 0 || len(r.stateSyncEvents.Replays) != 0 {
		t.Fatal("fresh boot counted historical events")
	}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 251, rec.fetch)
	if want := [][2]uint64{{251, 251}}; !slices.Equal(rec.ranges, want) {
		t.Errorf("ranges after bootstrap = %v, want %v", rec.ranges, want)
	}
	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{2}) {
		t.Errorf("commits after bootstrap = %v, want only next finalized event [2]", got)
	}
}

// A caught-up scan does no work; when a later finalized head contains no
// events, the successful empty scan still advances the cursor.
func TestScanStateSyncLogs_BootstrapAndEmptyScan(t *testing.T) {
	r := newStateSyncProvider(1000)
	rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{}}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec.fetch)
	if len(rec.ranges) != 0 {
		t.Fatalf("fresh boot fetched ranges %v, want none", rec.ranges)
	}
	if r.stateSyncCursor != 250 {
		t.Errorf("cursor = %d, want 250", r.stateSyncCursor)
	}
	if len(r.stateSyncEvents.Commits) != 0 || len(r.stateSyncEvents.Replays) != 0 {
		t.Errorf("fresh boot produced events")
	}

	// A second cycle with the same finalized head does nothing.
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec.fetch)
	if len(rec.ranges) != 0 {
		t.Errorf("caught-up scan refetched %v", rec.ranges)
	}

	// The next finalized block is empty, but it is still a successful scan.
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 251, rec.fetch)
	if want := [][2]uint64{{251, 251}}; !slices.Equal(rec.ranges, want) {
		t.Errorf("empty next-block ranges = %v, want %v", rec.ranges, want)
	}
	if r.stateSyncCursor != 251 {
		t.Errorf("cursor after empty scan = %d, want 251", r.stateSyncCursor)
	}
	if len(r.stateSyncEvents.Commits) != 0 || len(r.stateSyncEvents.Replays) != 0 {
		t.Errorf("empty scan produced events")
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

	// Bootstrap performs no backfill. The backlog from 300 through 550 is
	// therefore ongoing-process catch-up, which must retain the retry and
	// non-skipping cursor guarantees exercised below.
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 299, fetch)
	if len(ranges) != 0 {
		t.Fatalf("bootstrap fetched historical ranges %v", ranges)
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

// A bounded fetcher that re-delivers in-range logs (each response contains
// its logs twice, and each chunk's logs overlap the dedup keys committed by
// earlier chunks) must not double count. Out-of-range re-delivery is not a
// valid overlap: it rejects the chunk (see the rejection tests below).
func TestScanStateSyncLogs_OverlappingDeliveryDeduplicates(t *testing.T) {
	r := newStateSyncProvider(1000)
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 0

	allLogs := map[uint64][]types.Log{
		50:  {commitLog(50, 1, 0, true)},
		150: {commitLog(150, 2, 0, false)},
		180: {replayLog(180, 1, 0)},
	}

	fetch := func(_ context.Context, from, to uint64) ([]types.Log, error) {
		var logs []types.Log
		for block, ls := range allLogs {
			if block >= from && block <= to {
				// The same log delivered twice within one response.
				logs = append(logs, ls...)
				logs = append(logs, ls...)
			}
		}
		return logs, nil
	}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 200, fetch)

	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1, 2}) {
		t.Errorf("commit ids = %v, want [1 2] (in-range overlaps deduplicated)", got)
	}
	if got := replayIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1}) {
		t.Errorf("replay ids = %v, want [1] (in-range overlaps deduplicated)", got)
	}
	if r.stateSyncCursor != 200 {
		t.Errorf("cursor = %d, want 200", r.stateSyncCursor)
	}
}

// A response that fails validation must not advance the cursor: the next
// cycle retries the same range, and a corrected response is then counted.
// This is the difference between "drop the malformed event and continue"
// (which permanently skipped its blocks) and treating the chunk as
// unsuccessful.
func TestScanStateSyncLogs_InvalidChunkIsRetriedNotSkipped(t *testing.T) {
	r := newStateSyncProvider(100)
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 0

	calls := 0
	fetch := func(_ context.Context, from, to uint64) ([]types.Log, error) {
		calls++
		l := commitLog(100, 1, 0, true)
		if calls == 1 {
			// First response carries a non-canonical boolean (value 2).
			l.Data = append(make([]byte, 31), 2)
		}
		return []types.Log{l}, nil
	}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, fetch)
	if calls != 1 {
		t.Fatalf("fetch calls = %d, want 1", calls)
	}
	if r.stateSyncCursor != 0 {
		t.Fatalf("cursor advanced past an invalid chunk: %d, want 0", r.stateSyncCursor)
	}
	if len(r.stateSyncEvents.Commits) != 0 {
		t.Fatalf("invalid chunk counted %d commits, want 0", len(r.stateSyncEvents.Commits))
	}

	r.stateSyncEvents = &observer.StateSyncEvents{}
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, fetch)
	if calls != 2 {
		t.Fatalf("invalid range was not retried: fetch calls = %d, want 2", calls)
	}
	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1}) {
		t.Errorf("corrected response commits = %v, want [1]", got)
	}
	if r.stateSyncCursor != 100 {
		t.Errorf("cursor after corrected retry = %d, want 100", r.stateSyncCursor)
	}
}

// A malformed log reusing the (block hash, log index) key of a valid or an
// already-counted log must not slip past validation through the dedup
// path: the whole chunk is rejected, nothing is counted, the cursor does
// not advance, and the range is retried. Only a well-formed duplicate is
// suppressed as already counted.
func TestScanStateSyncLogs_MalformedDuplicateRejectsChunk(t *testing.T) {
	for _, prior := range []bool{false, true} {
		name := "duplicate within one response"
		if prior {
			name = "duplicate of an already-counted log"
		}
		t.Run(name, func(t *testing.T) {
			r := newStateSyncProvider(1000)
			r.stateSyncBootstrapped = true
			r.stateSyncCursor = 0
			good := commitLog(100, 1, 0, true)

			if prior {
				// Count the good log in a previous cycle so its key is
				// already in the dedup set when the malformed twin arrives.
				staged, ok := r.decodeStateSyncChunk(stateReceiverAddress, 100, 100, []types.Log{good})
				if !ok {
					t.Fatal("control response rejected")
				}
				r.commitStateSyncChunk(staged)
			}

			bad := good
			bad.Data = append(make([]byte, 31), 2)
			response := []types.Log{good, bad}

			calls := 0
			fetch := func(_ context.Context, _, _ uint64) ([]types.Log, error) {
				calls++
				return response, nil
			}

			r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, fetch)
			if r.stateSyncCursor != 0 {
				t.Fatalf("cursor advanced past a malformed duplicate: %d, want 0", r.stateSyncCursor)
			}
			if len(r.stateSyncEvents.Commits) != 0 {
				t.Fatalf("malformed duplicate counted commits: %d, want 0", len(r.stateSyncEvents.Commits))
			}

			// The corrected response must be fetched and accepted.
			response = []types.Log{good}
			r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, fetch)
			if calls != 2 {
				t.Fatalf("corrected response not retried: calls = %d, want 2", calls)
			}
			if r.stateSyncCursor != 100 {
				t.Fatalf("cursor after corrected retry = %d, want 100", r.stateSyncCursor)
			}
			// In the already-counted variant the valid log was counted
			// before the malformed twin arrived, so the corrected response
			// legitimately stages nothing; in the within-response variant
			// it stages the one valid log, exactly once.
			want := []uint64{1}
			if prior {
				want = nil
			}
			if got := commitIds(r.stateSyncEvents); !slices.Equal(got, want) {
				t.Errorf("commits after corrected retry = %v, want %v", got, want)
			}
		})
	}
}

// Inconsistent responses — a log above the finalized upper bound, a log with
// a zero block hash, a removed log — are rejected before counting and
// retried; the corrected response is then counted once.
func TestScanStateSyncLogs_InconsistentResponsesRejectedAndRetried(t *testing.T) {
	tests := []struct {
		name string
		bad  func() types.Log
	}{
		{
			name: "log above the finalized bound",
			bad:  func() types.Log { return commitLog(101, 1, 0, true) },
		},
		{
			name: "zero block hash",
			bad: func() types.Log {
				l := commitLog(100, 1, 0, true)
				l.BlockHash = common.Hash{}
				return l
			},
		},
		{
			name: "removed log",
			bad: func() types.Log {
				l := commitLog(100, 1, 0, true)
				l.Removed = true
				return l
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newStateSyncProvider(100)
			r.stateSyncBootstrapped = true
			r.stateSyncCursor = 0
			bad := true
			fetch := func(_ context.Context, from, to uint64) ([]types.Log, error) {
				if bad {
					return []types.Log{tt.bad()}, nil
				}
				return []types.Log{commitLog(100, 1, 0, true)}, nil
			}

			r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, fetch)
			if r.stateSyncCursor != 0 {
				t.Fatalf("cursor advanced past an inconsistent chunk: %d, want 0", r.stateSyncCursor)
			}
			if len(r.stateSyncEvents.Commits) != 0 {
				t.Fatalf("inconsistent chunk counted %d commits, want 0", len(r.stateSyncEvents.Commits))
			}

			bad = false
			r.stateSyncEvents = &observer.StateSyncEvents{}
			r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, fetch)
			if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1}) {
				t.Errorf("corrected response commits = %v, want [1]", got)
			}
			if r.stateSyncCursor != 100 {
				t.Errorf("cursor after corrected retry = %d, want 100", r.stateSyncCursor)
			}
		})
	}
}

// A state-sync record delivered to a receiver contract with no code emits no
// log at all: the record is absent from the commit denominator while
// lastStateId advances. This fixture models the documented absence with a
// log-free block between two committed records (the synthetic receipt itself
// is not part of these log fixtures, only its events are).
func TestScanStateSyncLogs_NoCodeReceiverRecordAbsent(t *testing.T) {
	r := newStateSyncProvider(100)
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 0
	rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{
		{1, 100}: {
			commitLog(10, 1, 0, true),
			// State ID 2 was delivered to a receiver with no code: Bor
			// records the sync but no StateCommitted event exists to count,
			// so there is no log fixture for it at block 50.
			commitLog(90, 3, 0, true),
		},
	}}

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, rec.fetch)

	if got := commitIds(r.stateSyncEvents); !slices.Equal(got, []uint64{1, 3}) {
		t.Errorf("commit ids = %v, want [1 3] (no-code record 2 emits nothing)", got)
	}
}

// Cancellation must stop the scan at the last completed chunk: the fetcher
// here cancels the context on its first call, so the first chunk completes
// and the scan yields before fetching a second one. The next cycle resumes.
func TestScanStateSyncLogs_CancellationYields(t *testing.T) {
	r := newStateSyncProvider(100)
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 100

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	calls := 0
	fetch := func(_ context.Context, from, to uint64) ([]types.Log, error) {
		calls++
		cancel()
		return nil, nil
	}

	r.scanStateSyncLogs(ctx, stateReceiverAddress, 100100, fetch)

	if calls > 1 {
		t.Fatalf("scan performed %d fetches after cancellation instead of yielding", calls)
	}
	if r.stateSyncCursor != 200 {
		t.Errorf("cursor = %d, want 200 (the chunk in flight when cancellation hit still completes)", r.stateSyncCursor)
	}
}

// A large catch-up backlog must not consume the whole cycle: the scan stops
// at its per-cycle chunk budget with the cursor at the last completed chunk,
// and the next cycle resumes exactly there.
func TestScanStateSyncLogs_PerCycleChunkBudget(t *testing.T) {
	r := newStateSyncProvider(100)
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 0

	var ranges [][2]uint64
	fetch := func(_ context.Context, from, to uint64) ([]types.Log, error) {
		ranges = append(ranges, [2]uint64{from, to})
		return nil, nil
	}

	// 1000 pending chunks after an outage.
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100000, fetch)

	if len(ranges) > stateSyncMaxChunksPerCycle {
		t.Fatalf("scan performed %d chunk fetches, want at most %d", len(ranges), stateSyncMaxChunksPerCycle)
	}
	wantCursor := uint64(stateSyncMaxChunksPerCycle * 100)
	if r.stateSyncCursor != wantCursor {
		t.Fatalf("cursor = %d, want %d (last completed chunk)", r.stateSyncCursor, wantCursor)
	}

	// The next cycle resumes at the last completed chunk.
	ranges = nil
	r.stateSyncEvents = &observer.StateSyncEvents{}
	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100000, fetch)
	if len(ranges) == 0 || ranges[0][0] != wantCursor+1 {
		t.Errorf("resumed at %v, want %d", ranges, wantCursor+1)
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
// records the current finalized head again. Neither process fetches or counts
// the historical lookback, and the new process's counters start at zero.
func TestScanStateSyncLogs_RestartSkipsHistoricalBackfill(t *testing.T) {
	newProviderWithHistory := func() (*RPCProvider, *rangeRecorder) {
		r := newStateSyncProvider(1000)
		rec := &rangeRecorder{logs: map[[2]uint64][]types.Log{
			{1, 100}:   {commitLog(50, 1, 0, false)},
			{201, 250}: {commitLog(250, 2, 0, false)},
		}}
		r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 250, rec.fetch)
		return r, rec
	}

	first, rec1 := newProviderWithHistory()
	second, rec2 := newProviderWithHistory()
	for i, rec := range []*rangeRecorder{rec1, rec2} {
		if len(rec.ranges) != 0 {
			t.Errorf("process %d fetched historical ranges %v", i+1, rec.ranges)
		}
	}
	if first.stateSyncCursor != 250 || second.stateSyncCursor != 250 {
		t.Errorf("restart cursors = %d, %d, want both 250", first.stateSyncCursor, second.stateSyncCursor)
	}
	if len(first.stateSyncEvents.Commits) != 0 || len(second.stateSyncEvents.Commits) != 0 {
		t.Error("restart counted historical events")
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

func TestNewRPCProvider_StateSyncEventsOptIn(t *testing.T) {
	interval := time.Second
	for _, enabled := range []bool{false, true} {
		r := NewRPCProvider(nil, observer.NewEventBus(), config.RPC{
			Name:            "test",
			URL:             "http://localhost",
			Label:           "test",
			Interval:        &interval,
			StateSyncEvents: enabled,
		})
		if r.stateSyncEventsEnabled != enabled {
			t.Errorf("stateSyncEventsEnabled = %v, want %v", r.stateSyncEventsEnabled, enabled)
		}
	}
}

// The per-RPC flag controls the actual eth_getLogs work, independently of the
// observer's metric registration. A configured receiver alone must not scan.
func TestRefreshStateSyncEvents_ProviderFlagControlsNetworkScan(t *testing.T) {
	var requests atomic.Int64
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var call struct {
			ID     json.RawMessage `json:"id"`
			Method string          `json:"method"`
		}
		if err := json.NewDecoder(req.Body).Decode(&call); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if call.Method != "eth_getLogs" {
			t.Errorf("unexpected RPC method %q", call.Method)
		}
		requests.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"jsonrpc": "2.0",
			"id":      call.ID,
			"result":  []types.Log{commitLog(100, 1, 0, true)},
		})
	}))
	defer server.Close()

	client, err := ethclient.Dial(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	for _, enabled := range []bool{false, true} {
		t.Run(map[bool]string{false: "disabled", true: "enabled"}[enabled], func(t *testing.T) {
			r := newStateSyncProvider(100)
			r.stateSyncEventsEnabled = enabled
			receiver := stateReceiverAddress.Hex()
			r.contracts.StateSyncReceiverAddress = &receiver
			r.stateSyncBootstrapped = true
			r.stateSyncCursor = 99
			r.finalizedHeight = 100

			before := requests.Load()
			r.refreshStateSyncEvents(context.Background(), client)
			gotRequests := requests.Load() - before
			wantRequests := int64(0)
			wantCursor := uint64(99)
			wantCommits := 0
			if enabled {
				wantRequests = 1
				wantCursor = 100
				wantCommits = 1
			}
			if gotRequests != wantRequests {
				t.Errorf("eth_getLogs requests = %d, want %d", gotRequests, wantRequests)
			}
			if r.stateSyncCursor != wantCursor {
				t.Errorf("cursor = %d, want %d", r.stateSyncCursor, wantCursor)
			}
			if got := len(r.stateSyncEvents.Commits); got != wantCommits {
				t.Errorf("decoded commits = %d, want %d", got, wantCommits)
			}
		})
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

// deliveryCountingObserver counts state-sync messages and the events they
// carry on a real event bus, which delivers messages from goroutines.
type deliveryCountingObserver struct {
	events   atomic.Int64
	messages atomic.Int64
	done     chan struct{}
	batches  chan *observer.StateSyncEvents
}

func (o *deliveryCountingObserver) Register(eb *observer.EventBus) {
	eb.Subscribe(topics.BorStateSyncEvents, o)
}

func (o *deliveryCountingObserver) GetCollectors() []prometheus.Collector {
	return nil
}

func (o *deliveryCountingObserver) Notify(_ context.Context, m observer.Message) {
	events := m.Data().(*observer.StateSyncEvents)
	o.events.Add(int64(len(events.Commits) + len(events.Replays)))
	o.messages.Add(1)
	if o.batches != nil {
		o.batches <- events
	}
	o.done <- struct{}{}
}

func waitForDelivery(t *testing.T, done <-chan struct{}) {
	t.Helper()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("event delivery timed out")
	}
}

// A failed refresh (here: an unreachable endpoint) must not re-deliver the
// previous cycle's events. The accumulation is reset before the fallible
// dial and publishing consumes the batch, so the post-failure publication
// carries an empty batch — not the stale events — and nothing double counts.
func TestRefreshStateDialFailureDoesNotRepublish(t *testing.T) {
	r := newStateSyncProvider(100)
	r.stateSyncEventsEnabled = true
	receiver := stateReceiverAddress.Hex()
	r.contracts.StateSyncReceiverAddress = &receiver
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 0
	r.bus = observer.NewEventBus()
	r.refreshStateTime = new(time.Duration)

	o := &deliveryCountingObserver{done: make(chan struct{}, 10)}
	r.bus.Subscribe(topics.BorStateSyncEvents, o)

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, func(context.Context, uint64, uint64) ([]types.Log, error) {
		return []types.Log{commitLog(100, 1, 0, false)}, nil
	})

	if err := r.PublishEvents(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitForDelivery(t, o.done)
	if got := o.events.Load(); got != 1 {
		t.Fatalf("events delivered = %d, want 1", got)
	}

	// A valid WebSocket URL whose endpoint is unreachable: DialContext fails
	// before the scan could reset the batch.
	r.url = "ws://127.0.0.1:1"
	if err := r.RefreshState(context.Background()); err == nil {
		t.Fatal("expected connection failure")
	}
	if err := r.PublishEvents(context.Background()); err != nil {
		t.Fatal(err)
	}

	waitForDelivery(t, o.done)
	if got := o.messages.Load(); got != 2 {
		t.Fatalf("messages delivered = %d, want 2 (one per publication)", got)
	}
	if got := o.events.Load(); got != 1 {
		t.Fatalf("stale events re-delivered after a failed refresh: %d events over 2 publications, want 1", got)
	}
}

// Publishing twice without a refresh in between must not re-deliver the
// consumed batch's events: the second publication carries an empty batch.
func TestPublishEventsConsumesBatch(t *testing.T) {
	r := newStateSyncProvider(100)
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 0
	r.bus = observer.NewEventBus()

	o := &deliveryCountingObserver{done: make(chan struct{}, 10)}
	r.bus.Subscribe(topics.BorStateSyncEvents, o)

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, func(context.Context, uint64, uint64) ([]types.Log, error) {
		return []types.Log{
			commitLog(100, 1, 0, true),
			replayLog(100, 2, 1),
		}, nil
	})

	if err := r.PublishEvents(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitForDelivery(t, o.done)
	if got := o.events.Load(); got != 2 {
		t.Fatalf("events delivered = %d, want 2 (one commit and one replay)", got)
	}

	if err := r.PublishEvents(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitForDelivery(t, o.done)
	if got := o.messages.Load(); got != 2 {
		t.Fatalf("messages delivered = %d, want 2 (one per publication)", got)
	}
	if got := o.events.Load(); got != 2 {
		t.Fatalf("republished batch re-delivered events: %d over 2 publications, want 2 (no double count)", got)
	}
}

// A published batch remains independent of later scans while the real bus
// delivers messages asynchronously. This is the ownership guarantee that lets
// the observer increment directly without batch digest or in-flight state.
func TestPublishEventsBatchIsolatedFromNextScan(t *testing.T) {
	r := newStateSyncProvider(100)
	r.stateSyncBootstrapped = true
	r.stateSyncCursor = 0
	r.bus = observer.NewEventBus()

	o := &deliveryCountingObserver{
		done:    make(chan struct{}, 2),
		batches: make(chan *observer.StateSyncEvents, 2),
	}
	r.bus.Subscribe(topics.BorStateSyncEvents, o)

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 100, func(context.Context, uint64, uint64) ([]types.Log, error) {
		return []types.Log{commitLog(100, 1, 0, true)}, nil
	})
	if err := r.PublishEvents(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitForDelivery(t, o.done)
	first := <-o.batches

	r.scanStateSyncLogs(context.Background(), stateReceiverAddress, 101, func(context.Context, uint64, uint64) ([]types.Log, error) {
		return []types.Log{commitLog(101, 2, 0, false)}, nil
	})
	if first == r.stateSyncEvents {
		t.Fatal("next scan reused the published batch pointer")
	}
	if got := commitIds(first); !slices.Equal(got, []uint64{1}) {
		t.Fatalf("next scan mutated the published batch: commits %v, want [1]", got)
	}

	if err := r.PublishEvents(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitForDelivery(t, o.done)
	second := <-o.batches
	if got := commitIds(second); !slices.Equal(got, []uint64{2}) {
		t.Errorf("second published batch commits = %v, want [2]", got)
	}
	if got := commitIds(first); !slices.Equal(got, []uint64{1}) {
		t.Errorf("first batch changed after second publication: %v, want [1]", got)
	}
}
