package media

import (
	"mime/multipart"
	"time"
)

type Media struct {
	ID            int64  `json:"id,omitempty"`
	Hash          string `json:"hash,omitempty"`
	File          *multipart.FileHeader
	URL           string    `json:"url,omitempty"`
	ThumbnailURL  string    `json:"thumbnail_url,omitempty"`
	TypeID        uint8     `json:"type_id,omitempty"`
	Extension     string    `json:"extension,omitempty"`
	IsBlacklisted bool      `json:"is_blacklisted,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

const (
	Image    = "img"
	Gif      = "gif"
	Video    = "vid"
	Document = "doc"
	Youtube  = "yt"
	// Vimeo = "vimeo"
	// Bitchute = "bitchute"
)
