package json_coder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Encode takes a value object and encodes it into JSON, and writes the encoded value to the given
// http response writer
func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("unable to encode json: %w", err)
	}
	return nil
}

// Decode accepts an http request and decodes the request body into the specified type
func Decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("unable to decode json: %w", err)
	}
	return v, nil
}
