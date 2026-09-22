package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/0xPolygon/panoptichain/api"
	"github.com/0xPolygon/panoptichain/observer"
)

// newFeeBalanceServer serves the Cosmos bank balances endpoint from a
// signer -> amount map. A signer absent from the map gets an empty balance
// list, which is what Heimdall returns for an account that has never been
// funded.
func newFeeBalanceServer(t *testing.T, amounts map[string]string, hits *atomic.Int64) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hits != nil {
			hits.Add(1)
		}

		const prefix = "/cosmos/bank/v1beta1/balances/"
		if !strings.HasPrefix(r.URL.Path, prefix) {
			http.NotFound(w, r)
			return
		}

		signer := strings.TrimPrefix(r.URL.Path, prefix)
		amount, ok := amounts[signer]
		if !ok {
			fmt.Fprint(w, `{"balances":[],"pagination":{"next_key":null,"total":"0"}}`)
			return
		}

		fmt.Fprintf(w, `{"balances":[{"denom":"pol","amount":%q}],"pagination":{"next_key":null,"total":"1"}}`, amount)
	}))
}

func newFeeBalanceProvider(serverURL string, validators ...api.Validator) *HeimdallProvider {
	curr := make(observer.ValidatorMap, len(validators))
	for _, v := range validators {
		curr[v.ID] = v
	}

	return &HeimdallProvider{
		heimdallURL:        serverURL,
		logger:             NewLogger(nil, "test"),
		validatorSets:      &observer.HeimdallValidatorSets{Curr: curr},
		feeBalancesEnabled: true,
		feeBalanceInterval: time.Minute,
		feeBalanceTimeout:  5 * time.Second,
	}
}

// byValidator indexes a sweep result for assertions.
func byValidator(b *observer.HeimdallFeeBalances) map[uint64]observer.HeimdallFeeBalance {
	out := make(map[uint64]observer.HeimdallFeeBalance, len(b.Balances))
	for _, fb := range b.Balances {
		out[fb.ValidatorID] = fb
	}
	return out
}

func TestRefreshFeeBalances_ReadsEveryValidator(t *testing.T) {
	server := newFeeBalanceServer(t, map[string]string{
		"0x1111111111111111111111111111111111111111": "1989000000000000000",
		"0x2222222222222222222222222222222222222222": "42",
	}, nil)
	defer server.Close()

	h := newFeeBalanceProvider(server.URL,
		api.Validator{ID: 1, Signer: "0x1111111111111111111111111111111111111111"},
		api.Validator{ID: 2, Signer: "0x2222222222222222222222222222222222222222"},
	)

	h.refreshFeeBalances(context.Background())

	if h.feeBalances == nil {
		t.Fatal("expected a sweep result")
	}
	if h.feeBalances.Failed != 0 {
		t.Fatalf("Failed = %d, want 0", h.feeBalances.Failed)
	}

	got := byValidator(h.feeBalances)
	if len(got) != 2 {
		t.Fatalf("got %d balances, want 2", len(got))
	}
	if v := got[1].Amount.String(); v != "1989000000000000000" {
		t.Errorf("validator 1 amount = %s, want 1989000000000000000", v)
	}
	if got[1].Denom != "pol" {
		t.Errorf("validator 1 denom = %q, want pol", got[1].Denom)
	}
}

// An account Heimdall reports as empty is an exhausted fee account, which is
// the whole point of the metric. It has to become an explicit zero, not a
// missing series.
func TestRefreshFeeBalances_EmptyAccountIsZero(t *testing.T) {
	server := newFeeBalanceServer(t, nil, nil)
	defer server.Close()

	h := newFeeBalanceProvider(server.URL,
		api.Validator{ID: 7, Signer: "0x7777777777777777777777777777777777777777"},
	)

	h.refreshFeeBalances(context.Background())

	got := byValidator(h.feeBalances)
	fb, ok := got[7]
	if !ok {
		t.Fatal("expected a balance for the unfunded validator")
	}
	if fb.Amount.Sign() != 0 {
		t.Errorf("amount = %s, want 0", fb.Amount)
	}
	if fb.Denom != feeDenom {
		t.Errorf("denom = %q, want %q", fb.Denom, feeDenom)
	}
}

// The sweep runs on its own cadence, not the provider's polling interval: at a
// 5s interval and ~100 validators, sweeping every cycle would be 20 requests a
// second against the Heimdall API for a value that moves by a flat fee per
// transaction.
func TestRefreshFeeBalances_HonoursItsOwnInterval(t *testing.T) {
	var hits atomic.Int64
	server := newFeeBalanceServer(t, map[string]string{
		"0x1111111111111111111111111111111111111111": "1",
	}, &hits)
	defer server.Close()

	h := newFeeBalanceProvider(server.URL,
		api.Validator{ID: 1, Signer: "0x1111111111111111111111111111111111111111"},
	)

	h.refreshFeeBalances(context.Background())
	first := hits.Load()

	h.feeBalances = nil
	h.refreshFeeBalances(context.Background())

	if hits.Load() != first {
		t.Fatalf("second cycle issued requests: %d, want %d", hits.Load(), first)
	}
	if h.feeBalances != nil {
		t.Fatal("expected no result on a cycle the sweep was not due")
	}

	// Once due, it sweeps again.
	h.nextFeeBalanceSweep = time.Now().Add(-time.Second)
	h.refreshFeeBalances(context.Background())

	if hits.Load() <= first {
		t.Fatal("sweep did not run once it was due again")
	}
}

// A failure must be counted rather than silently reducing the sweep to the
// validators that happened to answer, so the observer can tell a departed
// validator from an unreachable one.
func TestRefreshFeeBalances_CountsFailures(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "0x2222222222222222222222222222222222222222") {
			http.Error(w, "boom", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, `{"balances":[{"denom":"pol","amount":"5"}]}`)
	}))
	defer server.Close()

	h := newFeeBalanceProvider(server.URL,
		api.Validator{ID: 1, Signer: "0x1111111111111111111111111111111111111111"},
		api.Validator{ID: 2, Signer: "0x2222222222222222222222222222222222222222"},
	)

	h.refreshFeeBalances(context.Background())

	if h.feeBalances.Failed != 1 {
		t.Fatalf("Failed = %d, want 1", h.feeBalances.Failed)
	}
	if len(h.feeBalances.Balances) != 1 {
		t.Fatalf("got %d balances, want 1", len(h.feeBalances.Balances))
	}
}

// The sweep is the last step of RefreshState and is one request per validator.
// Its own deadline is what stops a degraded Heimdall API from spending the
// cycle here; it must return at that deadline rather than run to completion.
func TestRefreshFeeBalances_BoundedByItsOwnTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprint(w, `{"balances":[{"denom":"pol","amount":"5"}]}`)
	}))
	defer server.Close()

	h := newFeeBalanceProvider(server.URL,
		api.Validator{ID: 1, Signer: "0x1111111111111111111111111111111111111111"},
	)
	h.feeBalanceTimeout = 100 * time.Millisecond

	start := time.Now()
	h.refreshFeeBalances(context.Background())
	elapsed := time.Since(start)

	if elapsed > time.Second {
		t.Fatalf("sweep took %v; expected it to stop at its timeout", elapsed)
	}
	if h.feeBalances.Failed != 1 {
		t.Fatalf("Failed = %d, want 1", h.feeBalances.Failed)
	}
}

// Without a validator set there is nothing to sweep, and the cadence must not
// advance -- otherwise one failed validator-set refresh costs a full interval
// of fee balance coverage.
func TestRefreshFeeBalances_NoValidatorSetRetriesNextCycle(t *testing.T) {
	var hits atomic.Int64
	server := newFeeBalanceServer(t, nil, &hits)
	defer server.Close()

	h := newFeeBalanceProvider(server.URL)

	h.refreshFeeBalances(context.Background())

	if !h.nextFeeBalanceSweep.IsZero() {
		t.Fatal("cadence advanced despite there being nothing to sweep")
	}
	if hits.Load() != 0 {
		t.Fatalf("issued %d requests with no validator set", hits.Load())
	}
	if h.feeBalances != nil {
		t.Fatal("expected no result")
	}
}

func TestRefreshFeeBalances_Disabled(t *testing.T) {
	var hits atomic.Int64
	server := newFeeBalanceServer(t, nil, &hits)
	defer server.Close()

	h := newFeeBalanceProvider(server.URL,
		api.Validator{ID: 1, Signer: "0x1111111111111111111111111111111111111111"},
	)
	h.feeBalancesEnabled = false

	h.refreshFeeBalances(context.Background())

	if hits.Load() != 0 || h.feeBalances != nil {
		t.Fatal("disabled sweep still ran")
	}
}
