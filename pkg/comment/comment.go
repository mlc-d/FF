package comment

import (
	"time"

	"gitlab.com/mlc-d/ff/pkg/media"
)

type Comment struct {
	ID        int64
	UserID    int64
	ThreadID  int64
	Tag       string
	Content   string
	Media     *media.Media
	IsOP      bool
	Color     Color
	UniqueID  string
	CreatedAt time.Time
}

type Comments []*Comment

type Color uint8

const (
	White Color = iota + 1
	Red
	Yellow
	Blue
	Green
	Pink
	Black
	Orange
)
