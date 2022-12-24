package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/mlc-d/ff/pkg/user"
)

func Login(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`cannot decode json payload`))
		return
	}
	id, role, err := userService.Login(u)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(err.Error()))
		return
	}
	t, err := authService.CreateToken(id, role)
	authCookie := &http.Cookie{
		Name:  "jwt",
		Value: string(t),
	}
	http.SetCookie(w, authCookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`logged in!`))
}
