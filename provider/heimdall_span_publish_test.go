package provider

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

// When the cycle deadline cuts the walk short, the gauges still get the newest
// span Heimdall reported, not the span the walk stopped on. Otherwise a slow
// cycle could leave span_end_block behind the chain and false-fire the span-end
// alert, the same symptom as 2026-10-29.
func TestPublishSpans_TruncatedWalkPublishesFetchedLatest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bor/spans/latest" {
			// The deadline hits while the walk fetches its first span.
			cancel()
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(observer.HeimdallSpanV2{Span: *newSpan(103, 1300, 1399)})
	}))
	defer server.Close()

	h, r := newRecordingProvider(t, server.URL, newSpan(100, 1000, 1099))
	h.latestSpan = h.spans.Curr

	if err := h.refreshSpan(ctx); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr.ID != 100 {
		t.Fatalf("setup: expected the walk to stop at 100, Curr is %d", h.spans.Curr.ID)
	}

	h.publishSpans(context.Background())
	u := r.next(t)
	if len(u.Pairs) != 0 {
		t.Errorf("expected no pairs from a walk that made no progress, got %d", len(u.Pairs))
	}
	if u.Latest == nil || u.Latest.ID != 103 {
		t.Errorf("Latest = %v, want the fetched latest span 103", u.Latest)
	}
}

// A lagging Heimdall backend answering /bor/spans/latest with an older span
// must not move the gauges back.
func TestPublishSpans_StaleLatestDoesNotRegress(t *testing.T) {
	latest := newSpan(103, 1300, 1399)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(observer.HeimdallSpanV2{Span: *latest})
	}))
	defer server.Close()

	h, r := newRecordingProvider(t, server.URL, newSpan(103, 1300, 1399))
	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	h.publishSpans(context.Background())
	r.next(t)

	// The next response comes from a backend still on span 101.
	latest = newSpan(101, 1100, 1199)
	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	h.publishSpans(context.Background())

	if u := r.next(t); u.Latest == nil || u.Latest.ID != 103 {
		t.Errorf("Latest = %v, want span 103 to hold", u.Latest)
	}
}
