package zex

import (
	"errors"
	"log"
	"net/http"
)

// ErrHandler is a type that holds an error handler
type ErrHandler func(err error) http.HandlerFunc

// HandlerFuncWithErr is a type that holds an error handler function
type HandlerFuncWithErr func(w http.ResponseWriter, r *http.Request) error

// DefaultErrHandler is a function that returns a default error handler
var DefaultErrHandler = func(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e *Error

		if errors.As(err, &e) {
			log.Printf("Error: %s: %v\n", e.Error(), e.Internal())
			http.Error(w, e.Error(), e.Status())
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// WithError is a type that holds an error handler
type WithError struct {
	e ErrHandler
}

// NewWithError is a function that returns a new WithError
func NewWithErrorConverter(e ...ErrHandler) func(HandlerFuncWithErr) http.HandlerFunc {
	var handler ErrHandler
	if len(e) > 0 {
		handler = e[0]
	} else {
		handler = DefaultErrHandler
	}

	we := &WithError{e: handler}
	return we.Convert
}

// Convert is a method that converts a HandlerFuncWithErr to a http.HandlerFunc
func (w *WithError) Convert(h HandlerFuncWithErr) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := h(rw, r); err != nil {
			w.e(err)(rw, r)
		}
	}
}
