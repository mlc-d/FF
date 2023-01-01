package api

import (
	"encoding/json"
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/ff/web/api/internal"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	u := new(dto.User)
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
