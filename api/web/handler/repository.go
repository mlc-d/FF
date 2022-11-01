package handler

import (
	thread_service "gitlab.com/mlc-d/ff/pkg/thread/service"
	user_service "gitlab.com/mlc-d/ff/pkg/user/service"
)

var (
	userService   = user_service.NewUserService()
	threadService = thread_service.NewThreadService()
)
