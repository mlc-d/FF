package media

import "time"

type Media struct {
	ID            int64
	Hash          string
	URL           string
	ThumbnailURL  string
	TypeID        uint8
	Extension     string
	IsBlacklisted bool
	CreatedAt     time.Time
}

const (
	Image   = "img"
	Gif     = "gif"
	Video   = "vid"
	Youtube = "yt"
	// Vimeo = "vimeo"
	// Bitchute = "bitchute"
)
