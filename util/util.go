package util

import (
	"context"
	"time"

	"github.com/0xPolygon/panoptichain/log"
)

// BlockFor pauses execution until the current time is rounded to the nearest
// multiple of the specified duration, or until the context is canceled.
func BlockFor(ctx context.Context, duration time.Duration) {
	now := time.Now()
	rounded := now.Add(duration / 2).Round(duration)

	log.Trace().Time("now", now).Time("until", rounded).Msg("Blocking")

	timer := time.NewTimer(time.Until(rounded))
	defer timer.Stop()

	select {
	case <-timer.C:
		return
	case <-ctx.Done():
		return
	}
}

// RefreshTimeout is the hard ceiling on a single provider refresh cycle — a
// recovery net for a hung upstream, not a scheduling bound. It is deliberately
// generous (interval*4, floored at 30s) so it only trips on a genuine stall; a
// cycle merely running slower than its interval is surfaced by the overrun
// warning instead, since cancelling mid-cycle can drop in-progress work.
func RefreshTimeout(interval time.Duration) time.Duration {
	const minTimeout = 30 * time.Second
	if t := interval * 4; t > minTimeout {
		return t
	}
	return minTimeout
}
