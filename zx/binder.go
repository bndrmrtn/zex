package zx

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Bind decodes a JSON request into a value.
func Bind(r *http.Request, v any) error {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		return json.NewDecoder(r.Body).Decode(v)
	default:
		return errors.New("unsupported media type")
	}
}
