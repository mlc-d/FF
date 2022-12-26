package api

import (
	"encoding/json"
	"gitlab.com/mlc-d/ff/web/api/internal"
	"net/http"

	"gitlab.com/mlc-d/ff/pkg/user"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`cannot decode json payload`))
		return
	}
	id, _, err := internal.UserService.Register(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(id)
	w.WriteHeader(http.StatusCreated)
}
