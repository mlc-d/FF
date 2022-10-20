package web

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/mlc-d/ff/db"
	"gitlab.com/mlc-d/ff/pkg/entity"
	"gitlab.com/mlc-d/ff/pkg/entity/repository"
)

var (
	repo  repository.UserRepo = repository.NewUserRepo(sqlDB)
	sqlDB *sql.DB             = db.GetDB()
)

func registerRoutes(mux *chi.Mux) {
	mux.Get("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`test`))
	})
	mux.Post("/api/user", func(w http.ResponseWriter, r *http.Request) {
		user := new(entity.User)
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`cannot decode json payload`))
			return
		}
		id, err := repo.Register(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`cannot decode json payload`))
			return
		}
		err = json.NewEncoder(w).Encode(id)
		w.WriteHeader(http.StatusCreated)
	})
}
