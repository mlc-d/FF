package entity

import "time"

type Topic struct {
	ID        int64     `json:"id,omitempty"`
	ShortName string    `json:"short_name,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedBy int64     `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
