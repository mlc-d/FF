package thread

import (
	"time"

	"gitlab.com/mlc-d/ff/pkg/media"
)

type Thread struct {
	ID        int64
	Title     string
	Body      string
	MediaID   int64
	Media     *media.Media
	TopicID   int64
	UserID    int64
	CreatedAt time.Time
}
