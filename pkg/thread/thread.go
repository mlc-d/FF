package thread

import (
	"time"

	"gitlab.com/mlc-d/ff/pkg/comment"
	"gitlab.com/mlc-d/ff/pkg/media"
)

type Thread struct {
	ID        int64            `json:"id,omitempty"`
	TopicID   int64            `json:"topic_id,omitempty"`
	UserID    int64            `json:"user_id,omitempty"`
	Hash      string           `json:"hash,omitempty"`
	Title     string           `json:"title,omitempty"`
	Body      string           `json:"body,omitempty"`
	MediaID   int64            `json:"media_id,omitempty"`
	Media     *media.Media     `json:"media,omitempty"`
	Comments  comment.Comments `json:"comments,omitempty"`
	Sticky    bool             `json:"sticky,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
}

func New() *Thread {
	m := new(media.Media)
	return &Thread{
		Media: m,
	}
}
