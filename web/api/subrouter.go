package api

import (
	"github.com/go-chi/chi/v5"
	"gitlab.com/mlc-d/ff/web/api/protected"
)

var (
	router = chi.NewRouter()
)

func Router() *chi.Mux {
	router.Post("/login", Login)
	router.Post("/signup", SignUp)

	router.Mount("/protected", protected.Router())

	return router
}
