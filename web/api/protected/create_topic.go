package protected

import (
	"encoding/json"
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/go-jam"
	"net/http"
)

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	t := dto.NewTopic()

	err := r.ParseMultipartForm(4194304)
	_, file, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal([]byte(r.FormValue("topic")), t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid data"))
		return
	}

	_, claims, err := jam.FromContext(r.Context())

	userID := int64(claims["id"].(float64))
	t.CreatedBy = userID

	t.Media.File = file

	id, err := topicService.Create(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(id)
}
