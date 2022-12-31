package internal

import (
	"gitlab.com/mlc-d/ff/dto"
	"time"
)

type Comment struct {
	ID        int64      `json:"id,omitempty"`
	UserID    int64      `json:"user_id,omitempty"`
	ThreadID  int64      `json:"thread_id,omitempty"`
	Tag       string     `json:"tag,omitempty"`
	Content   string     `json:"content,omitempty"`
	MediaID   int64      `json:"media_id,omitempty"`
	Media     *dto.Media `json:"media,omitempty"`
	IsOP      bool       `json:"is_op,omitempty"`
	Color     Color      `json:"color,omitempty"`
	UniqueID  string     `json:"unique_id,omitempty"`
	Pinned    bool       `json:"pinned,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
}

type Comments []*Comment

type Color uint8
