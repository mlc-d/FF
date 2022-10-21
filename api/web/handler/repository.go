package handler

import (
	user_service "gitlab.com/mlc-d/ff/pkg/user/service"
)

var (
	userService = user_service.NewUserService()
)
