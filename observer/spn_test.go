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
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "1234567890", OnDemandGas: "42"},
		Requester:    requester,
		Hour:         "2026-08-06T12:00:00Z",
	})

	labels := []string{testNetwork.GetName(), "test", requester}

	if got, _ := gaugeValue(t, o.reserved, labels...); got != 1234567890 {
		t.Fatalf("reserved = %v, want 1234567890", got)
	}
	if got, _ := gaugeValue(t, o.on_demand, labels...); got != 42 {
		t.Fatalf("on_demand = %v, want 42", got)
	}
}

// A malformed value must leave the previous reading in place. Resetting to zero
// would render as "no gas used" on a dashboard rather than "no data".
func TestRequesterUsageObserver_MalformedValueKeepsLastReading(t *testing.T) {
	o := newUsageObserver(t)

	requester := "0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
	labels := []string{testNetwork.GetName(), "test", requester}

	notifyUsage(t, o, &UsageSummary{
		UsageSummary: &spnpb.UsageSummary{ReservedGas: "100", OnDemandGas: "200"},
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
}

// The provider leaves usage nil when the hour is empty or the call failed; the
// observer must not panic or publish anything for that.
func TestRequesterUsageObserver_IgnoresNilUsage(t *testing.T) {
	o := newUsageObserver(t)

	notifyUsage(t, o, nil)
	notifyUsage(t, o, &UsageSummary{Requester: "0xabc"})
}
