package handler

import (
	"gitlab.com/mlc-d/ff/pkg/auth"
	comment "gitlab.com/mlc-d/ff/pkg/comment/service"
	thread_service "gitlab.com/mlc-d/ff/pkg/thread/service"
	user_service "gitlab.com/mlc-d/ff/pkg/user/service"
)

var (
	userService    = user_service.NewUserService()
	threadService  = thread_service.NewThreadService()
	authService    = auth.NewJWTService()
	commentService = comment.NewCommentService()
)
