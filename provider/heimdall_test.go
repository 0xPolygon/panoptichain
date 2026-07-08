package provider

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/0xPolygon/panoptichain/config"
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
		maxSpanLag:  config.DefaultMaxSpanLag,
	}
}

// newProviderWithMaxLag creates a HeimdallProvider with custom maxSpanLag.
func newProviderWithMaxLag(serverURL string, curr *observer.HeimdallSpan, maxLag uint64) *HeimdallProvider {
	return &HeimdallProvider{
		heimdallURL: serverURL,
		spans:       &observer.HeimdallSpans{Curr: curr},
		logger:      NewLogger(nil, "test"),
		maxSpanLag:  maxLag,
	}
}

func TestRefreshSpan_Bootstrap(t *testing.T) {
	span := newSpan(100, 1000, 1099)
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": span,
	})
	defer server.Close()

	h := newProvider(server.URL, nil)

	if err := h.refreshSpan(context.Background()); err != nil {
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

	if err := h.refreshSpan(context.Background()); err != nil {
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

	if err := h.refreshSpan(context.Background()); err != nil {
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

	if err := h.refreshSpan(context.Background()); err != nil {
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
	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr.ID != 101 {
		t.Errorf("first call: expected span ID 101, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev.ID != 100 {
		t.Errorf("first call: expected prev span ID 100, got %d", h.spans.Prev.ID)
	}

	// Second call: should advance from 101 to 102
	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr.ID != 102 {
		t.Errorf("second call: expected span ID 102, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev.ID != 101 {
		t.Errorf("second call: expected prev span ID 101, got %d", h.spans.Prev.ID)
	}

	// Third call: should advance from 102 to 103
	if err := h.refreshSpan(context.Background()); err != nil {
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
	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}
	if h.spans.Curr != prevCurr {
		t.Error("fourth call: expected no change when already at latest span")
	}
}

func TestFetchSpan_RejectsZeroValueSpan(t *testing.T) {
	// Server returns a zero-value span (simulating error envelope decoding to zero struct)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(observer.HeimdallSpanV2{
			Span: observer.HeimdallSpan{ID: 0, StartBlock: 0, EndBlock: 0},
		})
	}))
	defer server.Close()

	h := newProvider(server.URL, nil)
	_, err := h.fetchSpan(context.Background(), "latest")

	if err == nil {
		t.Fatal("expected error for zero-value span, got nil")
	}

	if !errors.Is(err, ErrInvalidSpan) {
		t.Errorf("expected ErrInvalidSpan, got: %v", err)
	}
}

func TestFetchSpan_AcceptsValidSpanZero(t *testing.T) {
	// Span 0 is valid on mainnet; it has StartBlock=0 but EndBlock > 0
	span0 := newSpan(0, 0, 255)
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/0": span0,
	})
	defer server.Close()

	h := newProvider(server.URL, nil)
	span, err := h.fetchSpan(context.Background(), "0")

	if err != nil {
		t.Fatalf("expected no error for valid span 0, got: %v", err)
	}

	if span.ID != 0 {
		t.Errorf("expected span ID 0, got %d", span.ID)
	}
	if span.EndBlock != 255 {
		t.Errorf("expected EndBlock 255, got %d", span.EndBlock)
	}
}

func TestRefreshSpan_ExcessiveLagJumpsToLatest(t *testing.T) {
	// When lag exceeds maxSpanLag, should jump directly to latest
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": newSpan(120, 12000, 12099),
	})
	defer server.Close()

	// Current span is 100, latest is 120, lag = 20
	// With maxSpanLag = 5, should jump to latest
	h := newProviderWithMaxLag(server.URL, newSpan(100, 10000, 10099), 5)

	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	if h.spans.Curr.ID != 120 {
		t.Errorf("expected to jump to latest span 120, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev == nil || h.spans.Prev.ID != 100 {
		t.Errorf("expected Prev to be 100, got %v", h.spans.Prev)
	}
}

func TestRefreshSpan_WithinLagWalksSequentially(t *testing.T) {
	// When lag is within maxSpanLag, should walk sequentially
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": newSpan(105, 10500, 10599),
		"/bor/spans/101":    newSpan(101, 10100, 10199),
	})
	defer server.Close()

	// Current span is 100, latest is 105, lag = 5
	// With maxSpanLag = 10, should walk sequentially
	h := newProviderWithMaxLag(server.URL, newSpan(100, 10000, 10099), 10)

	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	// Should advance to 101, not jump to 105
	if h.spans.Curr.ID != 101 {
		t.Errorf("expected sequential walk to span 101, got %d", h.spans.Curr.ID)
	}
	if h.spans.Prev == nil || h.spans.Prev.ID != 100 {
		t.Errorf("expected Prev to be 100, got %v", h.spans.Prev)
	}
}

func TestRefreshSpan_ExactLagThresholdWalksSequentially(t *testing.T) {
	// When lag equals maxSpanLag exactly, should still walk sequentially
	server := newSpanServer(t, map[string]*observer.HeimdallSpan{
		"/bor/spans/latest": newSpan(110, 11000, 11099),
		"/bor/spans/101":    newSpan(101, 10100, 10199),
	})
	defer server.Close()

	// Current span is 100, latest is 110, lag = 10
	// With maxSpanLag = 10, should walk sequentially (lag is not > maxSpanLag)
	h := newProviderWithMaxLag(server.URL, newSpan(100, 10000, 10099), 10)

	if err := h.refreshSpan(context.Background()); err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	// Should advance to 101, not jump to 110
	if h.spans.Curr.ID != 101 {
		t.Errorf("expected sequential walk to span 101, got %d", h.spans.Curr.ID)
	}
}
