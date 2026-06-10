package api

import (
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

// GetJSON fetches JSON data from the specified URL and decodes it into the
// target variable.
func GetJSON(url string, target any) error {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cache-Control", "no-cache, no-store")
	req.Header.Set("Pragma", "no-cache")

	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

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
