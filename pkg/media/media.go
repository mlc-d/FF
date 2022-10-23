package media

type Media struct {
	ID           int64
	Hash         string
	URL          string
	ThumbnailURL string
	TypeID       uint8
	Extension    string
}

const (
	Image   = "img"
	Gif     = "gif"
	Video   = "vid"
	Youtube = "yt"
	// Vimeo = "vimeo"
	// Bitchute = "bitchute"
)
