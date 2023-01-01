package comment

import (
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/ff/pkg/comment/internal"
	"gitlab.com/mlc-d/ff/pkg/media"
	"gitlab.com/mlc-d/ff/pkg/thread"
	"time"
)

var (
	commentRepo   = internal.NewCommentRepo()
	mediaService  = media.NewService()
	threadService = thread.NewService()
)

// Service is the interface used by other packages to interact
// with the comments in the application
type Service interface {
	Post(c *dto.Comment) (*int64, error)
}

// service private struct implementing the interface Service
type service struct {
	repo  internal.CommentRepo
	media media.Service
}

// NewService returns a new [Service] interface
func NewService() Service {
	return &service{
		repo:  commentRepo,
		media: mediaService,
	}
}

// Post makes the proper validations on a [dto.Comment] instance,
// saves the [dto.Comment.Media] and creates the corresponding
// record in the database.
func (cs *service) Post(c *dto.Comment) (*int64, error) {
	comment := new(internal.Comment)

	comment.ID = c.ID
	comment.UserID = c.UserID
	comment.ThreadID = c.ThreadID
	comment.Tag = c.Tag
	comment.Content = c.Content
	comment.MediaID = c.MediaID

	comment.IsOP = threadService.GetUserID(c.ThreadID) == comment.UserID
	comment.CreatedAt = time.Now().UTC()
	comment.Color = cs.pickColor()
	if c.Media.File != nil {
		// TODO: handle this error
		_, _ = cs.media.Upload(c.Media)
	}
	return cs.repo.Post(comment)
}

// pickColor returns a number between 1 and 8, which are at the
// same time abstractions to represent colors
func (cs *service) pickColor() internal.Color {
	// FIXME: quick way to get a number between 1 and 8. Some colors should
	// be easier to obtain, so a new method must be implemented to get some
	// sort of hierarchy.
	x := time.Now().Unix() % 8
	return internal.Color(x + 1)
}
