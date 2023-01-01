package internal

import (
	"gitlab.com/mlc-d/ff/dto"
	"time"
)

type Thread struct {
	ID        int64
	TopicID   int64
	UserID    int64
	Hash      string
	Title     string
	Body      string
	MediaID   int64
	Media     *dto.Media
	Comments  *dto.Comments
	Sticky    bool
	CreatedAt time.Time
}
