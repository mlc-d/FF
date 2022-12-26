package protected

import (
	"github.com/go-chi/chi/v5"
	"gitlab.com/mlc-d/ff/web/api/internal"
	"gitlab.com/mlc-d/go-jam"
)

var (
	router = chi.NewRouter()
)

func Router() *chi.Mux {
	router.Use(jam.Verifier(internal.JWTService))
	router.Use(jam.Authenticator)
	return router
}
