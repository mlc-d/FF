package auth

import user_service "gitlab.com/mlc-d/ff/pkg/user/service"

type Service interface {
	Authenticate(name, password string) error
}

var (
	userService = user_service.NewUserService()
)

type authService struct {
	userService user_service.UserService
}

func NewAuthService() Service {
	return &authService{
		userService: userService,
	}
}

func (as *authService) Authenticate(name, password string) error {
	return nil
}
