package observer

import (
	"sync"
	"testing"

	spnpb "github.com/0xPolygon/panoptichain/proto/network"
)

// EventBus.Publish delivers each message on its own goroutine, so Notify runs
// concurrently. Exercise that shape directly: many goroutines delivering the
// same hour for the same requester must leave the counter at exactly one hour's
// gas, and must not race on the dedupe map.
func TestRequesterUsageObserver_ConcurrentNotifyIsSafe(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x8888888888888888888888888888888888888888"
	labels := []string{testNetwork.GetName(), "test", requester, "", "true"}

	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.Notify(nil, NewMessage(testNetwork, "test", &UsageSummary{
				UsageSummary: &spnpb.UsageSummary{ReservedGas: "1000000000", OnDemandGas: "0", TotalGas: "1000000000"},
				Requester:    requester,
				Billed:       true,
				Hour:         "2026-08-31T10:00:00Z",
			}))
		}()
	}
	wg.Wait()

	if got := counterValue(t, o.consumed, labels...); got != 1e9 {
		t.Fatalf("consumed = %v, want exactly one hour 1e9", got)
	}
}
