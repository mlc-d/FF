package internal

import (
	"mime/multipart"
	"time"
)

type Media struct {
	ID            int64
	Hash          string
	File          *multipart.FileHeader
	URL           string
	ThumbnailURL  string
	TypeID        uint8
	Extension     string
	IsBlacklisted bool
	CreatedAt     time.Time
}
