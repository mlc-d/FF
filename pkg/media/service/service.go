package media_service

import (
	"mime/multipart"

	"gitlab.com/mlc-d/ff/db"
	media_repo "gitlab.com/mlc-d/ff/pkg/media/repo"
)

var (
	sqlDB     = db.GetDB()
	mediaRepo = media_repo.NewMediaRepo(sqlDB)
)

type MediaService interface {
	Upload(file *multipart.FileHeader) (string, error)
}

type mediaService struct {
	repo media_repo.MediaRepo
}

func NewMediaService() MediaService {
	return &mediaService{
		repo: mediaRepo,
	}
}

func (ms mediaService) Upload(file *multipart.FileHeader) (string, error) {
	return ms.uploadFile(file)
}
