package internal

import (
	"gitlab.com/mlc-d/ff/dto"
	"time"
)

type Thread struct {
	ID        int64         `json:"id,omitempty"`
	TopicID   int64         `json:"topic_id,omitempty"`
	UserID    int64         `json:"user_id,omitempty"`
	Hash      string        `json:"hash,omitempty"`
	Title     string        `json:"title,omitempty"`
	Body      string        `json:"body,omitempty"`
	MediaID   int64         `json:"media_id,omitempty"`
	Media     *dto.Media    `json:"media,omitempty"`
	Comments  *dto.Comments `json:"comments,omitempty"`
	Sticky    bool          `json:"sticky,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
}

func New() *Thread {
	m := new(dto.Media)
	return &Thread{
		Media: m,
	}
}
