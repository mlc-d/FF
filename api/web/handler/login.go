package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/mlc-d/ff/pkg/user"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := new(user.User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`cannot decode json payload`))
		return
	}
	err = userService.Login(user)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`logged in!`))
}