package observer

import (
	"context"
	"os"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
	spnpb "github.com/0xPolygon/panoptichain/proto/network"
)

func gaugeValue(t *testing.T, g *prometheus.GaugeVec, labels ...string) (float64, bool) {
	t.Helper()

	m, err := g.GetMetricWithLabelValues(labels...)
	if err != nil {
		t.Fatalf("get metric: %v", err)
	}

	var out dto.Metric
	if err := m.Write(&out); err != nil {
		t.Fatalf("write metric: %v", err)
	}

	return out.GetGauge().GetValue(), out.Gauge != nil
}

func counterValue(t *testing.T, c *prometheus.CounterVec, labels ...string) float64 {
	t.Helper()

	m, err := c.GetMetricWithLabelValues(labels...)
	if err != nil {
		t.Fatalf("get metric: %v", err)
	}

	var out dto.Metric
	if err := m.Write(&out); err != nil {
		t.Fatalf("write metric: %v", err)
	}

	return out.GetCounter().GetValue()
}

// hasSeries reports whether a gauge holds the given series. Delete answers that
// without creating it, which GetMetricWith* would; and the collectors are
// registered on the process-wide registry, so every test in this package shares
// one gauge and absence cannot be asserted by counting series.
func hasSeries(g *prometheus.GaugeVec, labels ...string) bool {
	return g.DeleteLabelValues(labels...)
}

// The SPN provider is not tied to a chain; any network stands in for the label.
var testNetwork = &network.PolygonMainnet

// Register (and the metrics constructors it calls) dereference the global
// config, so give the test a minimal one.
func newUsageObserver(t *testing.T) *RequesterUsageObserver {
	t.Helper()

	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}

	o := new(RequesterUsageObserver)
	o.Register(NewEventBus())
	return o
}

func notifyUsage(t *testing.T, o *RequesterUsageObserver, usage *UsageSummary) {
	t.Helper()
	o.Notify(context.Background(), NewMessage(testNetwork, "test", usage))
}

func TestRequesterUsageObserver_SetsGaugesFromStrings(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{
			ReservedGas: "1234567890",
			OnDemandGas: "42",
			TotalGas:    "1234567932",
		},
		Requester: requester,
		Tag:       "katana",
		Hour:      "2026-08-06T12:00:00Z",
	})

	labels := []string{testNetwork.GetName(), "test", requester, "katana"}

	if got, _ := gaugeValue(t, o.reserved, labels...); got != 1234567890 {
		t.Fatalf("reserved = %v, want 1234567890", got)
	}
	if got, _ := gaugeValue(t, o.on_demand, labels...); got != 42 {
		t.Fatalf("on_demand = %v, want 42", got)
	}
	if got, _ := gaugeValue(t, o.hourly, labels...); got != 1234567932 {
		t.Fatalf("total = %v, want 1234567932", got)
	}
}

// The total is a convenience field. When the network omits it, the components
// it bills on still have to produce a total and therefore a cost.
func TestRequesterUsageObserver_FallsBackToComponentSum(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0xafb1d2c26654c85f51f550c97f16699da0293dee"
	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "6000000000", OnDemandGas: "4000000000"},
		Requester:    requester,
	})

	labels := []string{testNetwork.GetName(), "test", requester, ""}

	if got, _ := gaugeValue(t, o.hourly, labels...); got != 10000000000 {
		t.Fatalf("total = %v, want the component sum 10000000000", got)
	}
}

// An hour billed entirely on demand, with no reserved gas, still has to produce
// a total: the two components are independent and either can be zero.
func TestRequesterUsageObserver_OnDemandOnlyHourIsReported(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0xcd35546f2a72e64c6e19a64372511fe2b121fb46"
	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "0", OnDemandGas: "1000000000", TotalGas: "1000000000"},
		Requester:    requester,
		Tag:          "devtools",
	})

	labels := []string{testNetwork.GetName(), "test", requester, "devtools"}

	if got, _ := gaugeValue(t, o.hourly, labels...); got != 1000000000 {
		t.Fatalf("total = %v, want 1000000000", got)
	}
}

// A malformed value must leave the previous reading in place. Resetting to zero
// would render as "no gas used" on a dashboard rather than "no data".
func TestRequesterUsageObserver_MalformedValueKeepsLastReading(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
	labels := []string{testNetwork.GetName(), "test", requester, ""}

	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "100", OnDemandGas: "200", TotalGas: "300"},
		Requester:    requester,
	})
	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "", OnDemandGas: "not-a-number"},
		Requester:    requester,
	})

	if got, _ := gaugeValue(t, o.reserved, labels...); got != 100 {
		t.Fatalf("reserved = %v, want the previous 100", got)
	}
	if got, _ := gaugeValue(t, o.on_demand, labels...); got != 200 {
		t.Fatalf("on_demand = %v, want the previous 200", got)
	}
	// Neither the reported total nor the component sum is usable, so the total
	// must hold its last reading too rather than dropping to zero.
	if got, _ := gaugeValue(t, o.hourly, labels...); got != 300 {
		t.Fatalf("total = %v, want the previous 300", got)
	}
}

// The provider leaves usage nil when the hour is empty or the call failed; the
// observer must not panic or publish anything for that.
func TestRequesterUsageObserver_IgnoresNilUsage(t *testing.T) {
	o := newUsageObserver(t)

	notifyUsage(t, o, nil)
	notifyUsage(t, o, &UsageSummary{Requester: "0xabc"})
}

// The counters advance only for a bucket the provider flagged as new. The
// provider owns that decision; the observer must honour it, because adding a
// repeat would overstate consumption permanently.
func TestRequesterUsageObserver_CountersFollowNewHourFlag(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x2222222222222222222222222222222222222222"
	labels := []string{testNetwork.GetName(), "test", requester, "acme"}

	send := func(gas string, newHour bool) {
		notifyUsage(t, o, &UsageSummary{
			UsageSummary: &spnpb.UsageSummary{ReservedGas: gas, OnDemandGas: "0", TotalGas: gas},
			Requester:    requester,
			Tag:          "acme",
			Hour:         "2026-08-31T10:00:00Z",
			NewHour:      newHour,
		})
	}

	// One new bucket, then the same bucket re-reported across later cycles.
	send("1000000000", true)
	send("1000000000", false)
	send("1000000000", false)

	if got := counterValue(t, o.consumed, labels...); got != 1e9 {
		t.Fatalf("consumed = %v, want 1e9 (repeats must not accumulate)", got)
	}

	// A newly flagged bucket advances both.
	send("2000000000", true)

	if got := counterValue(t, o.consumed, labels...); got != 3e9 {
		t.Fatalf("consumed = %v, want 3e9", got)
	}

	// The gauge tracks the newest reading rather than summing, on every message.
	if got, _ := gaugeValue(t, o.hourly, labels...); got != 2e9 {
		t.Fatalf("hourly gauge = %v, want 2e9", got)
	}
}

// Requesters are counted independently.
func TestRequesterUsageObserver_CountersArePerRequester(t *testing.T) {
	o := newUsageObserver(t)

	a := "0x4444444444444444444444444444444444444444"
	b := "0x5555555555555555555555555555555555555555"

	for _, r := range []string{a, b} {
		notifyUsage(t, o, &UsageSummary{
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1000000000", OnDemandGas: "0", TotalGas: "1000000000"},
			Requester:    r,
			Hour:         "2026-08-31T10:00:00Z",
			NewHour:      true,
		})
	}

	for _, r := range []string{a, b} {
		labels := []string{testNetwork.GetName(), "test", r, ""}
		if got := counterValue(t, o.consumed, labels...); got != 1e9 {
			t.Fatalf("requester %s: consumed = %v, want 1e9", r, got)
		}
	}
}
