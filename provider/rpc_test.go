package provider

import (
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/contracts"
	"github.com/0xPolygon/panoptichain/observer"
)

// newHeaderBlock builds a NewHeaderBlock event with the given header block ID.
func newHeaderBlock(id int64) *contracts.RootChainNewHeaderBlock {
	return &contracts.RootChainNewHeaderBlock{HeaderBlockId: big.NewInt(id)}
}

// runCheckpointCycle simulates one poll cycle of refreshCheckpoint's dedup
// logic: seenCheckpoint decides whether the checkpoint is new, and the new path
// records it unseen (mirroring the store at the end of refreshCheckpoint, minus
// the RPC lookups). It returns whether this cycle would have published a
// countable checkpoint (Seen == false), which is what the observer increments
// its counters on.
func runCheckpointCycle(r *RPCProvider, event *contracts.RootChainNewHeaderBlock) bool {
	if !r.seenCheckpoint(event) {
		r.prevCheckpointID = new(big.Int).Set(event.HeaderBlockId)
		r.checkpointSignatures[false] = &observer.CheckpointSignatures{
			Event: event,
			Seen:  false,
		}
	}
	return !r.checkpointSignatures[false].Seen
}

func newCheckpointProvider() *RPCProvider {
	return &RPCProvider{
		checkpointSignatures: make(map[bool]*observer.CheckpointSignatures),
		logger:               NewLogger(nil, "test"),
	}
}

// TestSeenCheckpoint_CountsOncePerCheckpoint is the regression test for the
// over-counting bug: a checkpoint is republished on every poll cycle, but its
// signatures must only be counted once. Each cycle produces a countable
// (unseen) entry only when a genuinely new checkpoint is detected.
func TestSeenCheckpoint_CountsOncePerCheckpoint(t *testing.T) {
	a, b := newHeaderBlock(100), newHeaderBlock(200)

	// A realistic sequence: checkpoint A appears, then several idle polls with
	// no event in the scan window, then a new checkpoint B, then B is re-found
	// once at a scan-window boundary before going idle again.
	cycles := []*contracts.RootChainNewHeaderBlock{
		a,   // new checkpoint A detected
		nil, // idle poll, A republished
		nil, // idle poll, A republished
		nil, // idle poll, A republished
		b,   // new checkpoint B detected
		b,   // B re-found in an overlapping window
		nil, // idle poll, B republished
	}

	r := newCheckpointProvider()
	countable := 0
	for i, event := range cycles {
		if runCheckpointCycle(r, event) {
			countable++
			t.Logf("cycle %d: countable checkpoint published", i)
		}
	}

	// Only the two distinct checkpoints (A and B) should ever be counted.
	if countable != 2 {
		t.Fatalf("expected 2 countable checkpoints, got %d", countable)
	}
}

func TestSeenCheckpoint(t *testing.T) {
	tests := []struct {
		name           string
		prevID         *big.Int
		stored         *observer.CheckpointSignatures
		event          *contracts.RootChainNewHeaderBlock
		wantSeen       bool // return value
		wantStoredSeen bool // Seen of the stored entry after the call
		hasStored      bool
	}{
		{
			name:     "new checkpoint with no prior state is not seen",
			event:    newHeaderBlock(100),
			wantSeen: false,
		},
		{
			name:     "new checkpoint id is not seen",
			prevID:   big.NewInt(100),
			event:    newHeaderBlock(200),
			wantSeen: false,
		},
		{
			name:           "same checkpoint id is seen and the unseen entry is flipped",
			prevID:         big.NewInt(100),
			stored:         &observer.CheckpointSignatures{Event: newHeaderBlock(100), Seen: false},
			event:          newHeaderBlock(100),
			wantSeen:       true,
			wantStoredSeen: true,
			hasStored:      true,
		},
		{
			name:           "no event republishes the last checkpoint as seen",
			prevID:         big.NewInt(100),
			stored:         &observer.CheckpointSignatures{Event: newHeaderBlock(100), Seen: false},
			event:          nil,
			wantSeen:       true,
			wantStoredSeen: true,
			hasStored:      true,
		},
		{
			name:           "already seen entry stays seen",
			prevID:         big.NewInt(100),
			stored:         &observer.CheckpointSignatures{Event: newHeaderBlock(100), Seen: true},
			event:          nil,
			wantSeen:       true,
			wantStoredSeen: true,
			hasStored:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newCheckpointProvider()
			r.prevCheckpointID = tt.prevID
			if tt.hasStored {
				r.checkpointSignatures[false] = tt.stored
			}

			got := r.seenCheckpoint(tt.event)
			if got != tt.wantSeen {
				t.Errorf("seenCheckpoint() = %v, want %v", got, tt.wantSeen)
			}

			if tt.hasStored {
				if seen := r.checkpointSignatures[false].Seen; seen != tt.wantStoredSeen {
					t.Errorf("stored entry Seen = %v, want %v", seen, tt.wantStoredSeen)
				}
			}
		})
	}
}

// TestSeenCheckpoint_FreshCopy verifies that marking the stored entry seen does
// not mutate the previously published pointer, which may still be in flight in
// an observer goroutine.
func TestSeenCheckpoint_FreshCopy(t *testing.T) {
	r := newCheckpointProvider()
	r.prevCheckpointID = big.NewInt(100)
	published := &observer.CheckpointSignatures{Event: newHeaderBlock(100), Seen: false}
	r.checkpointSignatures[false] = published

	r.seenCheckpoint(nil)

	if published.Seen {
		t.Error("previously published entry was mutated in place")
	}
	if !r.checkpointSignatures[false].Seen {
		t.Error("stored entry was not marked seen")
	}
}

// TestNewRPCProvider_AccountBalanceBatchSize verifies the configured batch size
// is sanitized: unset, zero, or oversized values fall back to the default so the
// batch loops can never spin forever (size 0) or wrap negative (size > MaxInt).
func TestNewRPCProvider_AccountBalanceBatchSize(t *testing.T) {
	def := int(config.DefaultAccountBalanceBatchSize)
	u := func(v uint64) *uint64 { return &v }
	interval := time.Second

	tests := []struct {
		name string
		cfg  *uint64
		want int
	}{
		{name: "unset uses default", cfg: nil, want: def},
		{name: "zero falls back to default", cfg: u(0), want: def},
		{name: "valid value is used", cfg: u(500), want: 500},
		{name: "value above MaxInt32 falls back to default", cfg: u(math.MaxInt32 + 1), want: def},
		{name: "MaxUint64 falls back to default", cfg: u(math.MaxUint64), want: def},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.RPC{Name: "test", URL: "http://localhost", Label: "test", Interval: &interval, AccountBalanceBatchSize: tt.cfg}
			r := NewRPCProvider(nil, observer.NewEventBus(), cfg)
			if r.accountBalanceBatchSize != tt.want {
				t.Errorf("accountBalanceBatchSize = %d, want %d", r.accountBalanceBatchSize, tt.want)
			}
		})
	}
}

// TestShouldTrackBalance covers the balance-tracking precedence: an inline
// per-account TrackBalances override wins, otherwise an account is tracked
// unless its tag is in the excludeBalanceTags set.
func TestShouldTrackBalance(t *testing.T) {
	trueVal, falseVal := true, false

	tests := []struct {
		name     string
		excluded []string
		account  config.Account
		want     bool
	}{
		{
			name:    "no exclusions tracks by default",
			account: config.Account{Tag: "relayer"},
			want:    true,
		},
		{
			name:     "excluded tag is skipped",
			excluded: []string{"relayer"},
			account:  config.Account{Tag: "relayer"},
			want:     false,
		},
		{
			name:     "non-excluded tag is tracked",
			excluded: []string{"relayer"},
			account:  config.Account{Tag: "sequencer"},
			want:     true,
		},
		{
			name:     "inline true overrides exclusion",
			excluded: []string{"relayer"},
			account:  config.Account{Tag: "relayer", TrackBalances: &trueVal},
			want:     true,
		},
		{
			name:    "inline false overrides default",
			account: config.Account{Tag: "sequencer", TrackBalances: &falseVal},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			excludeBalanceTags := make(map[string]struct{}, len(tt.excluded))
			for _, tag := range tt.excluded {
				excludeBalanceTags[tag] = struct{}{}
			}
			r := &RPCProvider{excludeBalanceTags: excludeBalanceTags}
			if got := r.shouldTrackBalance(tt.account); got != tt.want {
				t.Errorf("shouldTrackBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
