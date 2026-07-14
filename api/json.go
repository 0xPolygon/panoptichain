package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPError represents an HTTP response with a non-2xx status code.
type HTTPError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *HTTPError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("HTTP %d %s: %s", e.StatusCode, e.Status, e.Body)
	}
	return fmt.Sprintf("HTTP %d %s", e.StatusCode, e.Status)
}

// sharedClient is a single, pooled HTTP client reused across all requests.
// Providers hit the same hosts (e.g. the Heimdall/Tendermint APIs) many times
// per refresh cycle, so keeping connections alive avoids re-dialing on every
// call. The per-request Timeout is kept as a hard ceiling even when a caller's
// context grants a longer deadline.
var sharedClient = newSharedClient()

func newSharedClient() *http.Client {
	// Clone the default transport so we inherit proxy, TLS, and HTTP/2
	// defaults, then raise the idle-connection pool limits (the default
	// MaxIdleConnsPerHost of 2 would defeat pooling for our per-host bursts).
	// Fall back to a plain client if something replaced DefaultTransport with a
	// non-*http.Transport, rather than panicking at package init.
	dt, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return &http.Client{Timeout: 10 * time.Second}
	}

	transport := dt.Clone()
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 10
	transport.IdleConnTimeout = 90 * time.Second

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}
}

// HTTPClient returns the shared, pooled HTTP client. Callers that need to issue
// requests GetJSON does not cover (e.g. POSTs) should use this so they share
// the same connection pool.
func HTTPClient() *http.Client {
	return sharedClient
}

// GetJSON fetches JSON data from the specified URL and decodes it into the
// target variable. The request is bound to ctx, so cancelling or expiring the
// context aborts an in-flight request.
func GetJSON(ctx context.Context, url string, target any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cache-Control", "no-cache, no-store")
	req.Header.Set("Pragma", "no-cache")

	r, err := sharedClient.Do(req)
	if err != nil {
		return err
	}
	// Drain any unread bytes (a trailing newline the decoder leaves, or an
	// undecoded error body) before closing so the connection returns to the
	// idle pool instead of being discarded.
	defer func() {
		_, _ = io.Copy(io.Discard, r.Body)
		r.Body.Close()
	}()

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(r.Body, 256))
		return &HTTPError{
			StatusCode: r.StatusCode,
			Status:     r.Status,
			Body:       string(body),
		}
	}

	return json.NewDecoder(r.Body).Decode(target)
}
