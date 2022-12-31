package internal

import "time"

type Topic struct {
	ID             int64
	ShortName      string
	Name           string
	ThumbnailURL   string
	IsNSFW         bool
	MaximumThreads uint16
	CreatedBy      int64
	CreatedAt      time.Time
}
