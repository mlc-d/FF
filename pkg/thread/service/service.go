package thread_service

import (
	"time"

	"gitlab.com/mlc-d/ff/pkg/hash"
	media_service "gitlab.com/mlc-d/ff/pkg/media/service"
	"gitlab.com/mlc-d/ff/pkg/thread"
	thread_repo "gitlab.com/mlc-d/ff/pkg/thread/repo"
)

type ThreadService interface {
	Post(t *thread.Thread) (*int64, error)
}

var (
	repo  = thread_repo.NewThreadRepo()
	media = media_service.NewMediaService()
)

type threadService struct {
	repo  thread_repo.ThreadRepo
	media media_service.MediaService
}

func NewThreadService() ThreadService {
	return &threadService{
		repo:  repo,
		media: media,
	}
}

func (ts *threadService) Post(t *thread.Thread) (*int64, error) {
	t.Hash = hash.GenerateRandomString(20)
	mediaID, err := ts.media.Upload(t.Media)
	if err != nil {
		return nil, err
	}
	t.MediaID = *mediaID
	t.CreatedAt = time.Now()
	id, err := ts.repo.Post(t)
	if err != nil {
		return nil, err
	}
	return id, nil
}
