package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/mlc-d/ff/api/web/handler"
)

func registerRoutes(mux *chi.Mux) {
	mux.Get("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`test`))
	})
	mux.Post("/api/user", handler.RegisterUser)
}
