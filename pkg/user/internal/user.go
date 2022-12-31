package internal

import "time"

type User struct {
	ID        int64
	Nick      string
	Password  string
	RoleID    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}
