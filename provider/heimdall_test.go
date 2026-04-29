package provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/0xPolygon/panoptichain/observer"
)

// newSpan creates a HeimdallSpan with the given parameters.
func newSpan(id, start, end uint64) *observer.HeimdallSpan {
	return &observer.HeimdallSpan{
		ID:         id,
		StartBlock: start,
		EndBlock:   end,
	}
}

// newSpanServer creates a test server that responds with spans based on a path-to-span map.
func newSpanServer(t *testing.T, spans map[string]*observer.HeimdallSpan) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		span, ok := spans[r.URL.Path]
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(observer.HeimdallSpanV2{Span: *span})
	}))
}

// newProvider creates a HeimdallProvider with the given server URL and current span.
func newProvider(serverURL string, curr *observer.HeimdallSpan) *HeimdallProvider {
	return &HeimdallProvider{
		heimdallURL: serverURL,
		spans:       &observer.HeimdallSpans{Curr: curr},
		logger:      NewLogger(nil, "test"),
	}
}

func TestRefreshSpan_Bootstrap(t *testing.T) {
	span := newSpan(100, 1000, 1099)
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": span,
	})
	defer server.Close()

	h := newProvider(server.URL, nil)

	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	if h.spans.Curr == nil {
		t.Fatal("expected Curr span to be set")
	}
	if h.spans.Curr.ID != 100 {
		t.Errorf("expected span ID 100, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev != nil {
		t.Error("expected Prev span to be nil on bootstrap")
	}
}

func TestRefreshSpan_Sequential(t *testing.T) {
	span101 := newSpan(101, 1100, 1199)
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": span101,
		"/bor/spans/101":    span101,
	})
	defer server.Close()

	h := newProvider(server.URL, newSpan(100, 1000, 1099))

	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	if h.spans.Curr == nil {
		t.Fatal("expected Curr span to be set")
	}
	if h.spans.Curr.ID != 101 {
		t.Errorf("expected span ID 101, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev == nil {
		t.Fatal("expected Prev span to be set")
	}
	if h.spans.Prev.ID != 100 {
		t.Errorf("expected prev span ID 100, got %d", h.spans.Prev.ID)
	}
}

func TestRefreshSpan_NextSpanNotAvailable(t *testing.T) {
	span := newSpan(100, 1000, 1099)
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": span,
	})
	defer server.Close()

	h := newProvider(server.URL, newSpan(100, 1000, 1099))
	originalCurr := h.spans.Curr

	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	if h.spans.Curr != originalCurr {
		t.Error("expected Curr span to remain unchanged when no new span available")
	}
	if h.spans.Prev != nil {
		t.Error("expected Prev span to remain nil when no new span available")
	}
}

func TestRefreshSpan_DetectsOverlappingSpans(t *testing.T) {
	// Span 101 starts at 1050, which overlaps with span 100's end block (1099)
	span101 := newSpan(101, 1050, 1149)
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": span101,
		"/bor/spans/101":    span101,
	})
	defer server.Close()

	h := newProvider(server.URL, newSpan(100, 1000, 1099))

	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	if h.spans.Prev == nil || h.spans.Curr == nil {
		t.Fatal("expected both Prev and Curr to be set")
	}

	if h.spans.Curr.StartBlock > h.spans.Prev.EndBlock {
		t.Errorf("expected overlapping spans: curr.StartBlock (%d) <= prev.EndBlock (%d)",
			h.spans.Curr.StartBlock, h.spans.Prev.EndBlock)
	}
}

func TestRefreshSpan_GapFillsOneAtATime(t *testing.T) {
	// When latest span is ahead by multiple IDs, we should only advance one at a time
	// to ensure each span transition is published to the observer for overlap detection
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": newSpan(103, 1300, 1399),
		"/bor/spans/101":    newSpan(101, 1100, 1199),
		"/bor/spans/102":    newSpan(102, 1200, 1299),
		"/bor/spans/103":    newSpan(103, 1300, 1399),
	})
	defer server.Close()

	h := newProvider(server.URL, newSpan(100, 1000, 1099))

	// First call: should advance from 100 to 101 only
	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr.ID != 101 {
		t.Errorf("first call: expected span ID 101, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev.ID != 100 {
		t.Errorf("first call: expected prev span ID 100, got %d", h.spans.Prev.ID)
	}

	// Second call: should advance from 101 to 102
	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr.ID != 102 {
		t.Errorf("second call: expected span ID 102, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev.ID != 101 {
		t.Errorf("second call: expected prev span ID 101, got %d", h.spans.Prev.ID)
	}

	// Third call: should advance from 102 to 103
	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr.ID != 103 {
		t.Errorf("third call: expected span ID 103, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev.ID != 102 {
		t.Errorf("third call: expected prev span ID 102, got %d", h.spans.Prev.ID)
	}

	// Fourth call: should not advance (already at latest)
	prevCurr := h.spans.Curr
	if err := h.refreshSpan(); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr != prevCurr {
		t.Error("fourth call: expected no change when already at latest span")
	}
}
