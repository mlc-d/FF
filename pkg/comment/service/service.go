package comment_service

import (
	"time"

	"gitlab.com/mlc-d/ff/pkg/comment"
	comment_repo "gitlab.com/mlc-d/ff/pkg/comment/repo"
	media_service "gitlab.com/mlc-d/ff/pkg/media/service"
)

var (
	commentRepo = comment_repo.NewCommentRepo()
	mediaRepo   = media_service.NewMediaService()
)

type CommentService interface {
	Post(c *comment.Comment) (*int64, error)
}

type commentService struct {
	repo  comment_repo.CommentRepo
	media media_service.MediaService
}

func NewCommentService() CommentService {
	return &commentService{
		repo:  commentRepo,
		media: mediaRepo,
	}
}

func (cs *commentService) Post(c *comment.Comment) (*int64, error) {
	c.CreatedAt = time.Now().UTC()
	c.Color = cs.pickColor()
	if c.Media.File != nil {
		cs.media.Upload(c.Media)
	}
	return cs.repo.Post(c)
}

func (cs *commentService) pickColor() comment.Color {
	// FIXME: quick way to get a number between 1 and 8. Some colors should
	// be easier to obtain, so a new method must be implemented to get some
	// sort of hierarchy.
	x := time.Now().Unix() % 8
	return comment.Color(x + 1)
}
