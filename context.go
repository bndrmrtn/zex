package zex

import (
	"fmt"
	"net/http"
	"strconv"
)

// ContextKey is the type for context keys
type ContextKey string

const (
	// ContextParams is the key for the context params
	ContextParams ContextKey = "params"
)

// Param returns the value of the current route parameter from the request context
func Param(r *http.Request, key string) string {
	params, ok := r.Context().Value(ContextParams).(map[string]string)
	if !ok {
		return ""
	}

	val, ok := params[key]
	if !ok {
		return ""
	}

	return val
}

// ParamInt returns the value of the current route parameter from the request context as an integer
func ParamInt(r *http.Request, key string) (int, error) {
	val := Param(r, key)
	if val == "" {
		return 0, fmt.Errorf("parameter %s not found", key)
	}

	return strconv.Atoi(val)
}
