package protected

import (
	"encoding/json"
	"gitlab.com/mlc-d/ff/dto"
	"net/http"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	cm := dto.NewComment()

	err := r.ParseMultipartForm(4194304)
	_, file, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid data"))
		return
	}

	err = json.Unmarshal([]byte(r.FormValue("comment")), cm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid data"))
		return
	}

	cm.Media.File = file

	id, err := commentService.Post(cm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(id)
}
