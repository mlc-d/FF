package internal

import "time"

type Topic struct {
	ID             int64     `json:"id,omitempty"`
	ShortName      string    `json:"short_name,omitempty"`
	Name           string    `json:"name,omitempty"`
	ThumbnailURL   string    `json:"thumbnail_url"`
	IsNSFW         bool      `json:"is_nsfw"`
	MaximumThreads uint16    `json:"maximum_threads"`
	CreatedBy      int64     `json:"created_by,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
