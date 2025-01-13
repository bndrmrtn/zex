package zex

import (
	"net/http"
	"testing"

	"github.com/bndrmrtn/zex/zx"
)

func Test_RunApp(t *testing.T) {
	app := New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ID is " + zx.Param(r, "id")))
	})

	e := NewWithErrorConverter()

	app.Get("/error/{num}", e(func(w http.ResponseWriter, r *http.Request) error {
		_, err := zx.ParamInt(r, "num")
		if err != nil {
			return err
		}
		w.Write([]byte("ID is " + zx.Param(r, "num")))
		return nil
	}))

	app.Serve(":3000")
}
