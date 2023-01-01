package internal

import (
	"gitlab.com/mlc-d/ff/dto"
	"time"
)

type Topic struct {
	ID             int64
	ShortName      string
	Name           string
	MediaID        *int64
	Media          *dto.Media
	IsNSFW         bool
	MaximumThreads uint16
	CreatedBy      int64
	CreatedAt      time.Time
}
