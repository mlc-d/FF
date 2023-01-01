package topic

import (
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/ff/pkg/errs"
	"gitlab.com/mlc-d/ff/pkg/media"
	"gitlab.com/mlc-d/ff/pkg/topic/internal"
	"time"
)

type Service interface {
	Create(t *dto.Topic) (*int64, error)
}

var (
	tr = internal.NewRepo()
	ms = media.NewService()
)

type service struct {
	tr internal.Repo
	ms media.Service
}

func NewService() Service {
	return &service{
		tr: tr,
		ms: ms,
	}
}

func (ts *service) Create(t *dto.Topic) (*int64, error) {
	topic := new(internal.Topic)
	topic.Name = t.Name
	topic.ShortName = t.ShortName
	topic.IsNSFW = t.IsNSFW
	topic.MaximumThreads = t.MaximumThreads
	topic.CreatedBy = t.CreatedBy
	topic.CreatedAt = time.Now()

	if validateMaximumThreads(topic.MaximumThreads) {
		return nil, errs.ErrInvalidThreadsCap
	}

	mediaID, err := ts.ms.Upload(t.Media)
	if err != nil {
		return nil, err
	}

	topic.MediaID = mediaID

	id, err := ts.tr.Create(topic)
	return id, err
}

const (
	MaximumThreadsGlobal = 128
	MinimumThreadsGlobal = 12
)

func validateMaximumThreads(mt uint16) bool {
	return mt > MaximumThreadsGlobal || mt < MinimumThreadsGlobal
}
