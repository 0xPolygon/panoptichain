package provider

import (
	"strconv"
	"testing"

	"github.com/0xPolygon/panoptichain/observer"
)

func vals(spec ...[2]int64) []*observer.HeimdallValidator {
	out := make([]*observer.HeimdallValidator, 0, len(spec))
	for i, s := range spec {
		out = append(out, &observer.HeimdallValidator{
			Address:          string(rune('A' + i)),
			ProposerPriority: strconv.FormatInt(s[0], 10),
			VotingPower:      strconv.FormatInt(s[1], 10),
		})
	}
	return out
}

// The winner is argmax(priority + voting_power), NOT argmax(priority). B wins
// on voting power despite A holding the higher priority -- which is precisely
// the case refreshMissedBlockProposal gets wrong.
func TestCometBFTProposersUsesVotingPower(t *testing.T) {
	// priorities are shifted by their average (0) first, so they are used as-is.
	got, err := cometBFTProposers(vals([2]int64{10, 1}, [2]int64{0, 100}), 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got[0] != "B" {
		t.Fatalf("expected B (0+100) to beat A (10+1), got %q", got[0])
	}
}

// After winning, a validator is decremented by the total voting power, so it
// should not win again immediately.
func TestCometBFTProposersRotates(t *testing.T) {
	got, err := cometBFTProposers(vals([2]int64{0, 10}, [2]int64{0, 9}), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got[0] != "A" || got[1] != "B" {
		t.Fatalf("expected A then B, got %v", got)
	}
}

// Ties break on the lower address.
func TestCometBFTProposersTieBreak(t *testing.T) {
	got, err := cometBFTProposers(vals([2]int64{5, 10}, [2]int64{5, 10}), 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got[0] != "A" {
		t.Fatalf("expected the lower address A to win the tie, got %q", got[0])
	}
}

// A malformed priority must be an error, not a silently-zero priority that
// reorders the whole set.
func TestCometBFTProposersRejectsMalformed(t *testing.T) {
	v := vals([2]int64{0, 1})
	v[0].ProposerPriority = ""
	if _, err := cometBFTProposers(v, 0); err == nil {
		t.Fatal("expected an error for an unparseable proposer_priority")
	}
}

// rounds+1 winners are returned, so winners[:round] are the validators that
// were selected and failed to propose.
func TestCometBFTProposersReturnsOnePerRound(t *testing.T) {
	got, err := cometBFTProposers(vals([2]int64{0, 5}, [2]int64{0, 4}, [2]int64{0, 3}), 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 winners for rounds 0..2, got %d: %v", len(got), got)
	}
}
