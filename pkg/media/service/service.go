package media_service

import (
	"gitlab.com/mlc-d/ff/db"
	"gitlab.com/mlc-d/ff/pkg/media"
	media_repo "gitlab.com/mlc-d/ff/pkg/media/repo"
)

var (
	sqlDB     = db.GetDB()
	mediaRepo = media_repo.NewMediaRepo(sqlDB)
	// FIXME: read these values from a config file
	allowedFormats = []string{"png", "jpg", "jpeg", "gif", "mp4", "webm", "3gp", "pdf"}
)

type MediaService interface {
	Upload(m *media.Media) (*int64, error)
}

type mediaService struct {
	repo media_repo.MediaRepo
}

func NewMediaService() MediaService {
	return &mediaService{
		repo: mediaRepo,
	}
}

func (ms *mediaService) Upload(m *media.Media) (*int64, error) {
	file := m.File
	return ms.uploadFile(file)
}
