package observer

import (
	"context"
	"os"
	"testing"

	"github.com/0xPolygon/panoptichain/config"
)

func newRoundObserver(t *testing.T) *HeimdallMissedBlockProposalObserver {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}
	o := new(HeimdallMissedBlockProposalObserver)
	o.Register(NewEventBus())
	return o
}

// A round-0 block missed nothing, but the counter must still exist at zero so
// it can be rate()'d on a healthy chain.
func TestProposalRound_ZeroRoundStillEmitsCounter(t *testing.T) {
	o := newRoundObserver(t)
	o.Notify(context.Background(), NewMessage(testNetwork, "test",
		HeimdallMissedBlockProposal{100: {Round: 0}}))

	if got := counterValue(t, o.failedRounds, testNetwork.GetName(), "test"); got != 0 {
		t.Fatalf("failedRounds = %v, want 0", got)
	}
}

// The whole point: a failed round is counted AND attributed to the validator
// that was selected and did not propose.
func TestProposalRound_AttributesTheMiss(t *testing.T) {
	o := newRoundObserver(t)
	const signer = "4CA9FF871C7AA1E7B64E1EAE110835F68D6A0BD4"

	o.Notify(context.Background(), NewMessage(testNetwork, "test",
		HeimdallMissedBlockProposal{
			100: {Round: 0},
			101: {Round: 1, Missed: []string{signer}},
		}))

	if got := counterValue(t, o.failedRounds, testNetwork.GetName(), "test"); got != 1 {
		t.Fatalf("failedRounds = %v, want 1", got)
	}
	if got := counterValue(t, o.missedProposal,
		testNetwork.GetName(), "test", CanonicalAddress(signer)); got != 1 {
		t.Fatalf("missedProposal = %v, want 1", got)
	}
}
