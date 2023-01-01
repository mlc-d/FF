package api

import (
	"database/sql"
	"encoding/json"
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/ff/pkg/errs"
	"gitlab.com/mlc-d/ff/web/api/internal"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	u := new(dto.User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`cannot decode json payload`))
		return
	}
	id, role, err := internal.UserService.Login(u)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errs.ErrInvalidCredentials
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, tokenString, err := internal.JWTService.Encode(map[string]any{
		"id":   id,
		"role": role,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	authCookie := &http.Cookie{
		Name:  "jwt",
		Value: tokenString,
	}
	http.SetCookie(w, authCookie)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`logged in!`))
}
