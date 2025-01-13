package zex

import (
	"net/http"
	"testing"
)

func Test_RunApp(t *testing.T) {
	app := New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ID is " + Param(r, "id")))
	})

	e := NewWithErrorConverter()

	app.Get("/error/{num}", e(func(w http.ResponseWriter, r *http.Request) error {
		_, err := ParamInt(r, "num")
		if err != nil {
			return err
		}
		w.Write([]byte("ID is " + Param(r, "num")))
		return nil
	}))

	app.Serve(":3000")
}
