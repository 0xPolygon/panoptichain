package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "test", "value": 42}`))
	}))
	defer server.Close()

	var result struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	err := GetJSON(server.URL, &result)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result.Name != "test" {
		t.Errorf("expected name 'test', got '%s'", result.Name)
	}
	if result.Value != 42 {
		t.Errorf("expected value 42, got %d", result.Value)
	}
}

func TestGetJSON_HTTPError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantBody   string
	}{
		{
			name:       "500 Internal Server Error",
			statusCode: http.StatusInternalServerError,
			body:       `{"code":13,"message":"span not found"}`,
			wantBody:   `{"code":13,"message":"span not found"}`,
		},
		{
			name:       "404 Not Found",
			statusCode: http.StatusNotFound,
			body:       "not found",
			wantBody:   "not found",
		},
		{
			name:       "503 Service Unavailable with empty body",
			statusCode: http.StatusServiceUnavailable,
			body:       "",
			wantBody:   "",
		},
		{
			name:       "400 Bad Request",
			statusCode: http.StatusBadRequest,
			body:       "invalid request",
			wantBody:   "invalid request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.body))
			}))
			defer server.Close()

			var result struct{}
			err := GetJSON(server.URL, &result)

			if err == nil {
				t.Fatal("expected error, got nil")
			}

			var httpErr *HTTPError
			if !errors.As(err, &httpErr) {
				t.Fatalf("expected HTTPError, got %T: %v", err, err)
			}

			if httpErr.StatusCode != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, httpErr.StatusCode)
			}

			if httpErr.Body != tt.wantBody {
				t.Errorf("expected body %q, got %q", tt.wantBody, httpErr.Body)
			}
		})
	}
}

func TestGetJSON_NetworkError(t *testing.T) {
	// Use an invalid URL to simulate a network error
	var result struct{}
	err := GetJSON("http://localhost:1", &result)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Should not be an HTTPError since the connection failed
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		t.Errorf("expected non-HTTPError, got HTTPError: %v", httpErr)
	}
}

func TestHTTPError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *HTTPError
		want string
	}{
		{
			name: "with body",
			err: &HTTPError{
				StatusCode: 500,
				Status:     "500 Internal Server Error",
				Body:       `{"error":"something failed"}`,
			},
			want: `HTTP 500 500 Internal Server Error: {"error":"something failed"}`,
		},
		{
			name: "without body",
			err: &HTTPError{
				StatusCode: 503,
				Status:     "503 Service Unavailable",
				Body:       "",
			},
			want: "HTTP 503 503 Service Unavailable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.want {
				t.Errorf("Error() = %q, want %q", got, tt.want)
			}
		})
	}
}
