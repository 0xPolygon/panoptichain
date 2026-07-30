package observer

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/0xPolygon/panoptichain/config"
	"github.com/0xPolygon/panoptichain/network"
)

// TestSensorBlocksObserverBoundsTheRange is the regression test for the observer
// wedge. Start and End derive from peer-announced block heights, and Notify loops
// once per height between them: measured at ~32ns per iteration, a single bogon
// announcement at 2^63 parked this goroutine roughly 9,000 years from returning.
// Both the provider (head-advance clamp) and this observer (range bound) now cut
// that off; this covers the observer half, which also protects the Datastore-backed
// provider the clamp does not.
func TestSensorBlocksObserverBoundsTheRange(t *testing.T) {
	// Register (and the metrics constructors it calls) dereference the global
	// config, so give the test a minimal one.
	dir := t.TempDir()
	if err := os.WriteFile(dir+"/config.yml", []byte("namespace: test\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	if err := config.Init(); err != nil {
		t.Fatalf("config.Init: %v", err)
	}

	o := &SensorBlocksObserver{}
	o.Register(NewEventBus())

	msg := NewMessage(&network.PolygonMainnet, "test", &SensorBlocks{
		Start: 91_000_000,
		End:   1 << 63, // one bogon announcement away
	})

	done := make(chan struct{})
	go func() {
		o.Notify(context.Background(), msg)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Notify did not return: the per-height loop is unbounded again")
	}
}
