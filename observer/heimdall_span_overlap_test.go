package observer

import "testing"

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
