package zx

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON response to the client.
func JSON(w http.ResponseWriter, code int, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(b)
	return err
}
