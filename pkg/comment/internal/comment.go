package internal

import (
	"gitlab.com/mlc-d/ff/dto"
	"time"
)

type Comment struct {
	ID        int64
	UserID    int64
	ThreadID  int64
	Tag       string
	Content   string
	MediaID   int64
	Media     *dto.Media
	IsOP      bool
	Color     Color
	UniqueID  string
	Pinned    bool
	CreatedAt time.Time
}

type Comments []*Comment

type Color uint8
