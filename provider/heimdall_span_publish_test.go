package provider

import (
	"context"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/0xPolygon/panoptichain/observer"
	"github.com/0xPolygon/panoptichain/observer/topics"
)

// spanRecorder captures every message published on topics.Span. The bus
// notifies on its own goroutines, so messages arrive on a channel.
type spanRecorder struct {
	got chan *observer.HeimdallSpanUpdates
}

func (r *spanRecorder) Register(eb *observer.EventBus) { eb.Subscribe(topics.Span, r) }

func (r *spanRecorder) Notify(_ context.Context, m observer.Message) {
	r.got <- m.Data().(*observer.HeimdallSpanUpdates)
}

func (r *spanRecorder) GetCollectors() []prometheus.Collector { return nil }

// next returns the one message a publishSpans call produced and fails if it
// produced none or more than one.
func (r *spanRecorder) next(t *testing.T) *observer.HeimdallSpanUpdates {
	t.Helper()

	var u *observer.HeimdallSpanUpdates
	select {
	case u = <-r.got:
	case <-time.After(time.Second):
		t.Fatal("expected a topics.Span message, got none")
	}

	select {
	case extra := <-r.got:
		t.Fatalf("expected exactly one topics.Span message, got another: %+v", extra)
	case <-time.After(50 * time.Millisecond):
	}

	return u
}

func newRecordingProvider(t *testing.T, serverURL string, curr *observer.HeimdallSpan) (*HeimdallProvider, *spanRecorder) {
	t.Helper()

	bus := observer.NewEventBus()
	r := &spanRecorder{got: make(chan *observer.HeimdallSpanUpdates, 8)}
	r.Register(bus)

	h := newProvider(serverURL, curr)
	h.bus = bus
	return h, r
}

// The 2026-10-29 rotation walk (15306 -> 15307 -> 15308) must reach the bus as
// a single message carrying both pairs in order. Published one message per
// pair, the bus could apply 15307 last and freeze the gauges on it.
func TestPublishSpans_WalkIsOneOrderedMessage(t *testing.T) {
	s15307 := newSpan(15307, 96396536, 96405055)
	s15308 := newSpan(15308, 96405056, 96411455)
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": s15308,
		"/bor/spans/15307":  s15307,
		"/bor/spans/15308":  s15308,
	})
	defer server.Close()

	h, r := newRecordingProvider(t, server.URL, newSpan(15306, 96405056, 96411455))

	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	h.publishSpans(context.Background())

	u := r.next(t)
	wantPrev := []uint64{15306, 15307}
	wantCurr := []uint64{15307, 15308}
	if len(u.Pairs) != len(wantCurr) {
		t.Fatalf("expected %d pairs, got %d", len(wantCurr), len(u.Pairs))
	}
	for i, p := range u.Pairs {
		if p.Prev == nil || p.Prev.ID != wantPrev[i] || p.Curr == nil || p.Curr.ID != wantCurr[i] {
			t.Errorf("pair %d: got Prev/Curr %v/%v, want %d/%d", i, p.Prev, p.Curr, wantPrev[i], wantCurr[i])
		}
	}
	if u.Latest == nil || u.Latest.ID != 15308 {
		t.Errorf("Latest = %v, want span 15308", u.Latest)
	}

	// An idle refresh still republishes the latest span, with no pairs, so a
	// stale gauge write is corrected without re-counting overlaps.
	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	h.publishSpans(context.Background())

	u = r.next(t)
	if len(u.Pairs) != 0 {
		t.Errorf("expected no pairs on an idle refresh, got %d", len(u.Pairs))
	}
	if u.Latest == nil || u.Latest.ID != 15308 {
		t.Errorf("Latest = %v, want span 15308", u.Latest)
	}
}

// Nothing is published before the first span is known.
func TestPublishSpans_NothingBeforeFirstSpan(t *testing.T) {
	h, r := newRecordingProvider(t, "", nil)
	h.publishSpans(context.Background())

	select {
	case u := <-r.got:
		t.Fatalf("expected no message before the first span, got %+v", u)
	case <-time.After(50 * time.Millisecond):
	}
}
