package observer

import (
	"context"
	"os"
	"testing"

	"github.com/0xPolygon/panoptichain/config"
)

func TestOverlappedBlocks(t *testing.T) {
	span := func(start, end uint64) *HeimdallSpan {
		return &HeimdallSpan{StartBlock: start, EndBlock: end}
	}

	for _, c := range []struct {
		name       string
		prev, curr *HeimdallSpan
		want       uint64
	}{
		// The ordinary case the old formula already got right.
		{"partial overlap", span(1000, 1099), span(1050, 1149), 50},
		{"adjacent", span(1000, 1099), span(1100, 1199), 0},
		{"identical", span(1000, 1099), span(1000, 1099), 100},
		// Mainnet 15302 -> 15303 (2026-10-29): the replacement starts before
		// 15302 did. The old prev.End - curr.Start + 1 gave 9139.
		{"replacement starts earlier", span(96398656, 96405055), span(96395917, 96405055), 6400},
		// Mainnet 15306 -> 15307: the replacement for the current range follows
		// the next-range proposal and shares no blocks with it. The old formula
		// gave 14920.
		{"replacement before prev", span(96405056, 96411455), span(96396536, 96405055), 0},
	} {
		if got := overlappedBlocks(c.prev, c.curr); got != c.want {
			t.Errorf("%s: overlappedBlocks = %d, want %d", c.name, got, c.want)
		}
	}
}

// newOverlapObserver sets up a registered HeimdallSpanObserver against a
// minimal config, like the other observer tests do.
func newOverlapObserver(t *testing.T) *HeimdallSpanObserver {
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

// Pins both counters through Notify. When an overlap is counted is unchanged
// by the intersection fix, so "PoS - Span overlap detected" (which fires on
// span_overlaps) behaves as before; only span_overlapped_blocks changes.
func TestSpanObserver_OverlapCounters(t *testing.T) {
	o := newOverlapObserver(t)

	span := func(id, start, end uint64) *HeimdallSpan {
		return &HeimdallSpan{ID: id, StartBlock: start, EndBlock: end}
	}

	for _, c := range []struct {
		name         string
		prev, curr   *HeimdallSpan
		wantOverlaps float64
		wantBlocks   float64
	}{
		{"partial overlap", span(100, 1000, 1099), span(101, 1050, 1149), 1, 50},
		{"adjacent", span(100, 1000, 1099), span(101, 1100, 1199), 0, 0},
		{"replacement starts earlier", span(15302, 96398656, 96405055), span(15303, 96395917, 96405055), 1, 6400},
		// Counted as an overlap, as before, but the spans share no blocks.
		// Whether this pair should count at all is an open question.
		{"replacement before prev", span(15306, 96405056, 96411455), span(15307, 96396536, 96405055), 1, 0},
	} {
		// The counters live on the process-wide registry, so give each case its
		// own provider label to read them independently.
		provider := "overlap-" + c.name
		o.Notify(context.Background(), NewMessage(testNetwork, provider, &HeimdallSpanUpdates{
			Pairs:  []*HeimdallSpans{{Prev: c.prev, Curr: c.curr}},
			Latest: c.curr,
		}))

		labels := []string{testNetwork.GetName(), provider}
		if got := counterValue(t, o.overlaps, labels...); got != c.wantOverlaps {
			t.Errorf("%s: span_overlaps = %.0f, want %.0f", c.name, got, c.wantOverlaps)
		}
		if got := counterValue(t, o.overlapped, labels...); got != c.wantBlocks {
			t.Errorf("%s: span_overlapped_blocks = %.0f, want %.0f", c.name, got, c.wantBlocks)
		}
	}
}
