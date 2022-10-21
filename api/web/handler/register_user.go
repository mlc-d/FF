package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/mlc-d/ff/pkg/user"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := new(user.User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`cannot decode json payload`))
		return
	}
	id, err := userService.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.NewEncoder(w).Encode(id)
	w.WriteHeader(http.StatusCreated)
}
