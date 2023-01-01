package thread

import (
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/ff/pkg/media"
	"gitlab.com/mlc-d/ff/pkg/thread/internal"
	"time"

	"gitlab.com/mlc-d/ff/pkg/hash"
)

type Service interface {
	Post(t *dto.Thread) (*int64, error)
	GetUserID(int64) int64
}

var (
	tr = internal.NewRepo()
	ms = media.NewService()
)

type service struct {
	repo  internal.Repo
	media media.Service
}

func NewService() Service {
	return &service{
		repo:  tr,
		media: ms,
	}
}

func (ts *service) Post(t *dto.Thread) (*int64, error) {

	thread := new(internal.Thread)
	thread.ID = t.ID
	thread.TopicID = t.TopicID
	thread.UserID = t.UserID
	thread.Title = t.Title
	thread.Body = t.Body
	thread.MediaID = t.MediaID
	thread.Sticky = t.Sticky

	t.Hash = hash.GenerateRandomString(20)
	mediaID, err := ts.media.Upload(t.Media)
	if err != nil {
		return nil, err
	}
	t.MediaID = *mediaID
	t.CreatedAt = time.Now()
	id, err := ts.repo.Post(thread)
	if err != nil {
		return nil, err
	}
	return id, nil
}
func (ts *service) GetUserID(threadID int64) int64 {
	uid := ts.repo.GetUserID(threadID)
	return *uid
}
