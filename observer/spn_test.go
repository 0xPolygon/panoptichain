package observer

import (
	"context"
	"math"
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

// The cost gauge is the total gas priced at the configured rate. It carries the
// currency as a label so a second contract's rate needs no second metric.
func TestRequesterUsageObserver_PricesTotalGas(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x5428abf0e5aec1be48597a984a4f9570d9236f29"
	notifyUsage(t, o, &UsageSummary{
		// Ten billion gas at $0.50 per billion is $5.00.
		UsageSummary:      &spnpb.UsageSummary{ReservedGas: "6000000000", OnDemandGas: "4000000000", TotalGas: "10000000000"},
		Requester:         requester,
		Tag:               "katana",
		RatePerBillionGas: 0.5,
		Currency:          "USD",
	})

	labels := []string{testNetwork.GetName(), "test", requester, "katana", "USD"}

	got, _ := gaugeValue(t, o.cost_hourly, labels...)
	if math.Abs(got-5.0) > 1e-9 {
		t.Fatalf("cost = %v, want 5.0", got)
	}
}

// Without pricing the cost series must stay absent. A zero would read as a free
// hour, which is a worse lie than no data.
func TestRequesterUsageObserver_NoPricingEmitsNoCost(t *testing.T) {
	o := newUsageObserver(t)

	// A requester unique to this test, so a cost series left behind by another
	// one cannot be mistaken for this observation's.
	requester := "0x1111111111111111111111111111111111111111"
	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "1", OnDemandGas: "2", TotalGas: "3"},
		Requester:    requester,
	})

	labels := []string{testNetwork.GetName(), "test", requester, ""}

	// The gas still lands; only the cost is withheld.
	if got, _ := gaugeValue(t, o.hourly, labels...); got != 3 {
		t.Fatalf("total = %v, want 3", got)
	}
	if hasSeries(o.cost_hourly, append(labels, "")...) {
		t.Fatal("expected no cost series without pricing")
	}
}

// The total is a convenience field. When the network omits it, the components
// it bills on still have to produce a total and therefore a cost.
func TestRequesterUsageObserver_FallsBackToComponentSum(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0xafb1d2c26654c85f51f550c97f16699da0293dee"
	notifyUsage(t, o, &UsageSummary{
		UsageSummary:      &spnpb.UsageSummary{ReservedGas: "6000000000", OnDemandGas: "4000000000"},
		Requester:         requester,
		RatePerBillionGas: 0.5,
		Currency:          "USD",
	})

	labels := []string{testNetwork.GetName(), "test", requester, ""}

	if got, _ := gaugeValue(t, o.hourly, labels...); got != 10000000000 {
		t.Fatalf("total = %v, want the component sum 10000000000", got)
	}
	if got, _ := gaugeValue(t, o.cost_hourly, append(labels, "USD")...); math.Abs(got-5.0) > 1e-9 {
		t.Fatalf("cost = %v, want 5.0", got)
	}
}

// An hour billed entirely on demand, with no reserved gas, still has to produce
// a total and a cost: the two components are independent and either can be zero.
func TestRequesterUsageObserver_OnDemandOnlyHourIsReported(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0xcd35546f2a72e64c6e19a64372511fe2b121fb46"
	notifyUsage(t, o, &UsageSummary{
		UsageSummary:      &spnpb.UsageSummary{ReservedGas: "0", OnDemandGas: "1000000000", TotalGas: "1000000000"},
		Requester:         requester,
		Tag:               "devtools",
		RatePerBillionGas: 0.5,
		Currency:          "USD",
	})

	labels := []string{testNetwork.GetName(), "test", requester, "devtools"}

	if got, _ := gaugeValue(t, o.hourly, labels...); got != 1000000000 {
		t.Fatalf("total = %v, want 1000000000", got)
	}
	if got, _ := gaugeValue(t, o.cost_hourly, append(labels, "USD")...); math.Abs(got-0.5) > 1e-9 {
		t.Fatalf("cost = %v, want 0.5", got)
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

// The counters exist to be range-summed, which only works if each hourly bucket
// is added exactly once. The provider re-reads the same bucket every poll, so
// repeated delivery of one hour must not advance them.
func TestRequesterUsageObserver_CountersAddEachHourOnce(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x2222222222222222222222222222222222222222"
	labels := []string{testNetwork.GetName(), "test", requester, "acme"}
	costLabels := append(append([]string{}, labels...), "USD")

	hour := func(h, gas string) *UsageSummary {
		return &UsageSummary{
			UsageSummary:      &spnpb.UsageSummary{ReservedGas: gas, OnDemandGas: "0", TotalGas: gas},
			Requester:         requester,
			Tag:               "acme",
			Hour:              h,
			RatePerBillionGas: 0.5,
			Currency:          "USD",
		}
	}

	// One hour, delivered three times as three polls would.
	for i := 0; i < 3; i++ {
		notifyUsage(t, o, hour("2026-08-31T10:00:00Z", "1000000000"))
	}

	if got := counterValue(t, o.consumed, labels...); got != 1e9 {
		t.Fatalf("after 3 polls of one hour, consumed = %v, want 1e9", got)
	}
	if got := counterValue(t, o.cost_total, costLabels...); math.Abs(got-0.5) > 1e-9 {
		t.Fatalf("cost_total = %v, want 0.5", got)
	}

	// A new hour advances both counters.
	notifyUsage(t, o, hour("2026-08-31T11:00:00Z", "2000000000"))

	if got := counterValue(t, o.consumed, labels...); got != 3e9 {
		t.Fatalf("after a second hour, consumed = %v, want 3e9", got)
	}
	if got := counterValue(t, o.cost_total, costLabels...); math.Abs(got-1.5) > 1e-9 {
		t.Fatalf("cost_total = %v, want 1.5", got)
	}

	// The gauge, unlike the counter, tracks the newest hour rather than summing.
	if got, _ := gaugeValue(t, o.hourly, labels...); got != 2e9 {
		t.Fatalf("hourly gauge = %v, want the newest hour 2e9", got)
	}
}

// An hour older than one already counted must not be added. The provider takes
// the newest bucket, but a late or reordered delivery must not double count.
func TestRequesterUsageObserver_CountersIgnoreOlderHour(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x3333333333333333333333333333333333333333"
	labels := []string{testNetwork.GetName(), "test", requester, ""}

	send := func(h, gas string) {
		notifyUsage(t, o, &UsageSummary{
			UsageSummary: &spnpb.UsageSummary{ReservedGas: gas, OnDemandGas: "0", TotalGas: gas},
			Requester:    requester,
			Hour:         h,
		})
	}

	send("2026-08-31T12:00:00Z", "5000000000")
	send("2026-08-31T09:00:00Z", "9000000000") // older, must be ignored

	if got := counterValue(t, o.consumed, labels...); got != 5e9 {
		t.Fatalf("consumed = %v, want 5e9 (older hour ignored)", got)
	}
}

// Requesters are counted independently: one advancing must not consume another's
// budget for the same hour.
func TestRequesterUsageObserver_CountersArePerRequester(t *testing.T) {
	o := newUsageObserver(t)

	a := "0x4444444444444444444444444444444444444444"
	b := "0x5555555555555555555555555555555555555555"

	for _, r := range []string{a, b} {
		notifyUsage(t, o, &UsageSummary{
			UsageSummary: &spnpb.UsageSummary{ReservedGas: "1000000000", OnDemandGas: "0", TotalGas: "1000000000"},
			Requester:    r,
			Hour:         "2026-08-31T10:00:00Z",
		})
	}

	for _, r := range []string{a, b} {
		labels := []string{testNetwork.GetName(), "test", r, ""}
		if got := counterValue(t, o.consumed, labels...); got != 1e9 {
			t.Fatalf("requester %s: consumed = %v, want 1e9", r, got)
		}
	}
}

// Without pricing the gas counter still advances; only the cost counter stays
// absent, so adding pricing later starts costing from that point rather than
// retroactively.
func TestRequesterUsageObserver_UnpricedAdvancesGasCounterOnly(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x6666666666666666666666666666666666666666"
	labels := []string{testNetwork.GetName(), "test", requester, ""}

	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "7000000000", OnDemandGas: "0", TotalGas: "7000000000"},
		Requester:    requester,
		Hour:         "2026-08-31T10:00:00Z",
	})

	if got := counterValue(t, o.consumed, labels...); got != 7e9 {
		t.Fatalf("consumed = %v, want 7e9", got)
	}
	if o.cost_total.DeleteLabelValues(append(append([]string{}, labels...), "")...) {
		t.Fatal("expected no cost_total series without pricing")
	}
}

// A bucket with no hour cannot be deduplicated, so the counters must not move.
// A flat counter is recoverable; an overstated one is not.
func TestRequesterUsageObserver_NoHourLeavesCountersAlone(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0x7777777777777777777777777777777777777777"
	labels := []string{testNetwork.GetName(), "test", requester, ""}

	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "4000000000", OnDemandGas: "0", TotalGas: "4000000000"},
		Requester:    requester,
	})

	// The gauge still reports: an hourless bucket is usable as a level, just not
	// as something to accumulate.
	if got, _ := gaugeValue(t, o.hourly, labels...); got != 4e9 {
		t.Fatalf("hourly = %v, want 4e9", got)
	}
	if o.consumed.DeleteLabelValues(labels...) {
		t.Fatal("expected no consumed series for a bucket with no hour")
	}
}
