package media

import (
	"gitlab.com/mlc-d/ff/db"
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/ff/pkg/media/internal"
)

var (
	sqlDB     = db.GetDB()
	mediaRepo = internal.NewMediaRepo(sqlDB)
	// FIXME: read these values from a config file
	allowedFormats = []string{"png", "jpg", "jpeg", "gif", "mp4", "webm", "3gp", "pdf"}
)

type Service interface {
	Upload(m *dto.Media) (*int64, error)
}

type service struct {
	repo internal.MediaRepo
}

func NewService() Service {
	return &service{
		repo: mediaRepo,
	}
}

func (ms *service) Upload(m *dto.Media) (*int64, error) {
	file := m.File
	return ms.uploadFile(file)
}
