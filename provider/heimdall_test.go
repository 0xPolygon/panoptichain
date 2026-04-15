package provider

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/0xPolygon/panoptichain/observer"
)

func TestRefreshSpan_Bootstrap(t *testing.T) {
	// Mock server returns latest span
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bor/spans/latest" {
			t.Errorf("unexpected path: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		resp := observer.HeimdallSpanV2{
			Span: observer.HeimdallSpan{
				ID:         100,
				StartBlock: 1000,
				EndBlock:   1099,
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	h := &HeimdallProvider{
		heimdallURL: server.URL,
		spans:       &observer.HeimdallSpans{},
		logger:      NewLogger(nil, "test"),
	}

	err := h.refreshSpan()
	if err != nil {
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
	// Mock server returns latest span first, then span by ID
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/bor/spans/latest":
			resp := observer.HeimdallSpanV2{
				Span: observer.HeimdallSpan{
					ID:         101,
					StartBlock: 1100,
					EndBlock:   1199,
				},
			}
			json.NewEncoder(w).Encode(resp)
		case "/bor/spans/101":
			resp := observer.HeimdallSpanV2{
				Span: observer.HeimdallSpan{
					ID:         101,
					StartBlock: 1100,
					EndBlock:   1199,
				},
			}
			json.NewEncoder(w).Encode(resp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	h := &HeimdallProvider{
		heimdallURL: server.URL,
		spans: &observer.HeimdallSpans{
			Curr: &observer.HeimdallSpan{
				ID:         100,
				StartBlock: 1000,
				EndBlock:   1099,
			},
		},
		logger: NewLogger(nil, "test"),
	}

	err := h.refreshSpan()
	if err != nil {
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
	// Mock server returns latest span with same ID as current (no new span)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/bor/spans/latest" {
			resp := observer.HeimdallSpanV2{
				Span: observer.HeimdallSpan{
					ID:         100,
					StartBlock: 1000,
					EndBlock:   1099,
				},
			}
			json.NewEncoder(w).Encode(resp)
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	h := &HeimdallProvider{
		heimdallURL: server.URL,
		spans: &observer.HeimdallSpans{
			Curr: &observer.HeimdallSpan{
				ID:         100,
				StartBlock: 1000,
				EndBlock:   1099,
			},
		},
		logger: NewLogger(nil, "test"),
	}

	originalCurr := h.spans.Curr

	err := h.refreshSpan()
	if err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	// Curr should remain unchanged
	if h.spans.Curr != originalCurr {
		t.Error("expected Curr span to remain unchanged when no new span available")
	}
	if h.spans.Prev != nil {
		t.Error("expected Prev span to remain nil when no new span available")
	}
}

func TestRefreshSpan_DetectsOverlappingSpans(t *testing.T) {
	// This test verifies that sequential fetching allows overlap detection
	// The actual overlap detection happens in the observer, but we verify
	// that Prev and Curr are set correctly for the observer to detect it

	// Mock server returns latest span first, then span by ID
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/bor/spans/latest":
			resp := observer.HeimdallSpanV2{
				Span: observer.HeimdallSpan{
					ID:         101,
					StartBlock: 1050, // Overlaps with prev.EndBlock (1099)
					EndBlock:   1149,
				},
			}
			json.NewEncoder(w).Encode(resp)
		case "/bor/spans/101":
			resp := observer.HeimdallSpanV2{
				Span: observer.HeimdallSpan{
					ID:         101,
					StartBlock: 1050, // Overlaps with prev.EndBlock (1099)
					EndBlock:   1149,
				},
			}
			json.NewEncoder(w).Encode(resp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	h := &HeimdallProvider{
		heimdallURL: server.URL,
		spans: &observer.HeimdallSpans{
			Curr: &observer.HeimdallSpan{
				ID:         100,
				StartBlock: 1000,
				EndBlock:   1099,
			},
		},
		logger: NewLogger(nil, "test"),
	}

	err := h.refreshSpan()
	if err != nil {
		t.Fatalf("refreshSpan() error: %v", err)
	}

	// Verify spans are set up for overlap detection
	if h.spans.Prev == nil || h.spans.Curr == nil {
		t.Fatal("expected both Prev and Curr to be set")
	}

	// Verify the overlap condition that observer checks
	if h.spans.Curr.StartBlock > h.spans.Prev.EndBlock {
		t.Errorf("expected overlapping spans: curr.StartBlock (%d) <= prev.EndBlock (%d)",
			h.spans.Curr.StartBlock, h.spans.Prev.EndBlock)
	}
}

func TestRefreshSpan_GapFillsOneAtATime(t *testing.T) {
	// When latest span is ahead by multiple IDs, we should only advance one at a time
	// to ensure each span transition is published to the observer for overlap detection
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spans := map[string]observer.HeimdallSpan{
			"/bor/spans/latest": {ID: 103, StartBlock: 1300, EndBlock: 1399},
			"/bor/spans/101":    {ID: 101, StartBlock: 1100, EndBlock: 1199},
			"/bor/spans/102":    {ID: 102, StartBlock: 1200, EndBlock: 1299},
			"/bor/spans/103":    {ID: 103, StartBlock: 1300, EndBlock: 1399},
		}
		span, ok := spans[r.URL.Path]
		if !ok {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(observer.HeimdallSpanV2{Span: span})
	}))
	defer server.Close()

	h := &HeimdallProvider{
		heimdallURL: server.URL,
		spans: &observer.HeimdallSpans{
			Curr: &observer.HeimdallSpan{
				ID:         100,
				StartBlock: 1000,
				EndBlock:   1099,
			},
		},
		logger: NewLogger(nil, "test"),
	}

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
