package observer

import (
	"context"
	"os"
	"testing"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/config"
)

func newSpanObserver(t *testing.T) *HeimdallSpanObserver {
	t.Helper()

	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}

	o := new(HeimdallSpanObserver)
	o.Register(NewEventBus())

	return o
}

func heimdallSpan(id, start, end, producer uint64) *HeimdallSpan {
	return &HeimdallSpan{
		ID:                id,
		StartBlock:        start,
		EndBlock:          end,
		SelectedProducers: []api.Validator{{ID: producer}},
	}
}

// The mainnet spans from the 2026-10-29 producer rotations. 15307 replaced the
// current range after 15306 had already proposed the next one, and 15308
// re-proposed it. One refresh walked 15306 -> 15307 -> 15308; published as two
// messages, the 15307 pair was applied last and froze span_end_block at
// 96405055, false-firing "Block height too close to span end" for two hours.
func TestSpanObserver_RotationWalkEndsOnLatest(t *testing.T) {
	o := newSpanObserver(t)
	const provider = "rotation"

	s15306 := heimdallSpan(15306, 96405056, 96411455, 93)
	s15307 := heimdallSpan(15307, 96396536, 96405055, 91)
	s15308 := heimdallSpan(15308, 96405056, 96411455, 93)

	o.Notify(context.Background(), NewMessage(testNetwork, provider, &HeimdallSpanUpdates{
		Pairs: []*HeimdallSpans{
			{Prev: s15306, Curr: s15307},
			{Prev: s15307, Curr: s15308},
		},
		Latest: s15308,
	}))

	labels := []string{testNetwork.GetName(), provider}
	assertSpanGauges(t, o, labels, s15308)

	// Only 15306 -> 15307 overlaps (15307 starts before 15306 ends); 15308
	// starts right after 15307 ends.
	if got := counterValue(t, o.overlaps, labels...); got != 1 {
		t.Errorf("span_overlaps = %.0f, want 1", got)
	}
}

// A stale message applied after a newer one (the bus gives no ordering between
// messages) is overwritten by the next refresh, which republishes the latest
// span with no pairs. The idle refresh must not count the overlap again.
func TestSpanObserver_StaleWriteCorrectedByNextRefresh(t *testing.T) {
	o := newSpanObserver(t)
	const provider = "stale"

	s15306 := heimdallSpan(15306, 96405056, 96411455, 93)
	s15307 := heimdallSpan(15307, 96396536, 96405055, 91)
	s15308 := heimdallSpan(15308, 96405056, 96411455, 93)

	notify := func(u *HeimdallSpanUpdates) {
		o.Notify(context.Background(), NewMessage(testNetwork, provider, u))
	}

	// The newer refresh lands first, then the older one overwrites it.
	notify(&HeimdallSpanUpdates{Pairs: []*HeimdallSpans{{Prev: s15307, Curr: s15308}}, Latest: s15308})
	notify(&HeimdallSpanUpdates{Pairs: []*HeimdallSpans{{Prev: s15306, Curr: s15307}}, Latest: s15307})

	labels := []string{testNetwork.GetName(), provider}
	if got := must(gaugeValue(t, o.endBlock, labels...)); got != 96405055 {
		t.Fatalf("setup: span_end_block = %.0f, want the stale 96405055", got)
	}

	// Next idle refresh: no new pairs, latest is still 15308.
	notify(&HeimdallSpanUpdates{Latest: s15308})
	assertSpanGauges(t, o, labels, s15308)

	if got := counterValue(t, o.overlaps, labels...); got != 1 {
		t.Errorf("span_overlaps = %.0f, want 1 (idle refresh must not re-count)", got)
	}
}

func assertSpanGauges(t *testing.T, o *HeimdallSpanObserver, labels []string, want *HeimdallSpan) {
	t.Helper()

	for _, c := range []struct {
		name string
		got  float64
		want float64
	}{
		{"span_id", must(gaugeValue(t, o.spanID, labels...)), float64(want.ID)},
		{"span_start_block", must(gaugeValue(t, o.startBlock, labels...)), float64(want.StartBlock)},
		{"span_end_block", must(gaugeValue(t, o.endBlock, labels...)), float64(want.EndBlock)},
		{"span_producer", must(gaugeValue(t, o.producer, labels...)), float64(want.SelectedProducers[0].ID)},
	} {
		if c.got != c.want {
			t.Errorf("%s = %.0f, want %.0f", c.name, c.got, c.want)
		}
	}
}

func must(v float64, ok bool) float64 {
	if !ok {
		return -1
	}
	return v
}
