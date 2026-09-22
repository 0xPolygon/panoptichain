package observer

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/0xPolygon/panoptichain/config"
)

func newFeeBalanceObserver(t *testing.T) *HeimdallValidatorFeeBalanceObserver {
	t.Helper()

	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}

	o := new(HeimdallValidatorFeeBalanceObserver)
	o.Register(NewEventBus())

	return o
}

func feeBalance(id uint64, signer, denom string, amount int64) HeimdallFeeBalance {
	return HeimdallFeeBalance{
		ValidatorID:   id,
		SignerAddress: signer,
		Denom:         denom,
		Amount:        big.NewInt(amount),
	}
}

// The balance lands on the gauge under the canonical address spelling, which is
// the one the rest of the heimdall_ family already uses.
func TestFeeBalance_RecordsBalance(t *testing.T) {
	o := newFeeBalanceObserver(t)

	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{
			feeBalance(193, "0x6297093F882D6ACE2E737AC81AC473100455A4D7", "pol", 1_989_000),
		},
	}))

	got, _ := gaugeValue(t, o.balance, testNetwork.GetName(), "test", "193",
		"0x6297093f882d6ace2e737ac81ac473100455a4d7", "pol")
	if got != 1_989_000 {
		t.Fatalf("balance = %v, want 1989000", got)
	}

	if swept, _ := gaugeValue(t, o.swept, testNetwork.GetName(), "test"); swept != 1 {
		t.Fatalf("swept = %v, want 1", swept)
	}
}

// A validator that leaves the set must lose its series. Left in place it would
// sit frozen at its last balance, and a threshold alert on it would page
// forever with nothing to act on.
func TestFeeBalance_RetiresDepartedValidator(t *testing.T) {
	o := newFeeBalanceObserver(t)
	const signer = "0xaaaa000000000000000000000000000000000001"

	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(1, signer, "pol", 5)},
	}))
	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(2, "0xaaaa000000000000000000000000000000000002", "pol", 7)},
	}))

	if hasSeries(o.balance, testNetwork.GetName(), "test", "1", signer, "pol") {
		t.Fatal("departed validator's series was not retired")
	}
}

// A sweep that did not reach every validator says nothing about the ones it
// missed, so it must not retire them.
func TestFeeBalance_PartialSweepKeepsSeries(t *testing.T) {
	o := newFeeBalanceObserver(t)
	const signer = "0xbbbb000000000000000000000000000000000001"

	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(1, signer, "pol", 5)},
	}))
	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(2, "0xbbbb000000000000000000000000000000000002", "pol", 7)},
		Failed:   1,
	}))

	if !hasSeries(o.balance, testNetwork.GetName(), "test", "1", signer, "pol") {
		t.Fatal("partial sweep retired a series it had no evidence about")
	}
	if failed, _ := gaugeValue(t, o.failed, testNetwork.GetName(), "test"); failed != 1 {
		t.Fatalf("failed = %v, want 1", failed)
	}
}

// The sweep that follows a partial one is complete, and must retire what the
// partial one was not allowed to.
func TestFeeBalance_CompleteSweepAfterPartialRetires(t *testing.T) {
	o := newFeeBalanceObserver(t)
	const signer = "0xcccc000000000000000000000000000000000001"

	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(1, signer, "pol", 5)},
	}))
	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(2, "0xcccc000000000000000000000000000000000002", "pol", 7)},
		Failed:   1,
	}))
	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(2, "0xcccc000000000000000000000000000000000002", "pol", 7)},
	}))

	if hasSeries(o.balance, testNetwork.GetName(), "test", "1", signer, "pol") {
		t.Fatal("complete sweep did not retire the stale series")
	}
}

// An exhausted fee account is the condition worth alerting on, so a zero must
// be published rather than dropped.
func TestFeeBalance_ZeroIsRecorded(t *testing.T) {
	o := newFeeBalanceObserver(t)
	const signer = "0xdddd000000000000000000000000000000000001"

	o.Notify(context.Background(), NewMessage(testNetwork, "test", &HeimdallFeeBalances{
		Balances: []HeimdallFeeBalance{feeBalance(9, signer, "pol", 0)},
	}))

	got, ok := gaugeValue(t, o.balance, testNetwork.GetName(), "test", "9", signer, "pol")
	if !ok || got != 0 {
		t.Fatalf("balance = %v (ok=%v), want 0", got, ok)
	}
}

// A nil payload must not panic the observer.
func TestFeeBalance_NilPayload(t *testing.T) {
	o := newFeeBalanceObserver(t)
	o.Notify(context.Background(), NewMessage(testNetwork, "test", (*HeimdallFeeBalances)(nil)))
}
