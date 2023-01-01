package dto

import "time"

type Topic struct {
	ID             int64     `json:"id,omitempty"`
	ShortName      string    `json:"short_name,omitempty"`
	Name           string    `json:"name,omitempty"`
	MediaID        *int64    `json:"media_id"`
	Media          *Media    `json:"media"`
	IsNSFW         bool      `json:"is_nsfw"`
	MaximumThreads uint16    `json:"maximum_threads"`
	CreatedBy      int64     `json:"created_by,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func NewTopic() *Topic {
	m := new(Media)
	return &Topic{
		Media: m,
	}
}

/*

{
	"short_name": "UFF",
	"name": "Random",
	"is_nsfw": true,
	"maximum_threads": 48
}

*/
