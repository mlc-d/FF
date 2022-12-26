package protected

import (
	"encoding/json"
	"gitlab.com/mlc-d/go-jam"
	"net/http"

	"gitlab.com/mlc-d/ff/pkg/thread"
)

func CreateThread(w http.ResponseWriter, r *http.Request) {
	t := thread.New()

	err := r.ParseMultipartForm(4194304)
	_, file, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid data"))
		return
	}

	err = json.Unmarshal([]byte(r.FormValue("thread")), t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid data"))
		return
	}

	_, claims, err := jam.FromContext(r.Context())
	userID := claims["id"].(int64)
	t.UserID = userID

	t.Media.File = file

	id, err := threadService.Post(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(id)
}
