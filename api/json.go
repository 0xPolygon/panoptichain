package api

import (
	"encoding/json"
	"net/http"
	"time"
)

// GetJSON fetches JSON data from the specified URL and decodes it into the
// target variable.
func GetJSON(url string, target any) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
