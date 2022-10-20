package user

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Password  string    `json:"password,omitempty"`
	RoleID    uint8     `json:"role_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
