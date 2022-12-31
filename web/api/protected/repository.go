package protected

import (
	"gitlab.com/mlc-d/ff/pkg/comment"
	thread_service "gitlab.com/mlc-d/ff/pkg/thread"
)

var (
	threadService  = thread_service.NewService()
	commentService = comment.NewService()
)
