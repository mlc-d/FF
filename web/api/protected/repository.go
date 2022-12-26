package protected

import (
	comment "gitlab.com/mlc-d/ff/pkg/comment/service"
	thread_service "gitlab.com/mlc-d/ff/pkg/thread/service"
)

var (
	threadService  = thread_service.NewThreadService()
	commentService = comment.NewCommentService()
)
