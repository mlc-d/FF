package user_service

import (
	"regexp"
	"strings"

	"gitlab.com/mlc-d/ff/db"
	"gitlab.com/mlc-d/ff/pkg/errs"
	"gitlab.com/mlc-d/ff/pkg/user"
	user_repo "gitlab.com/mlc-d/ff/pkg/user/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	Dev = iota
	Admin
	Mod
	Anon
	AllowedChars = `[\p{L}\p{N}]`
)

var (
	sqlDB    = db.GetDB()
	userRepo = user_repo.NewUserRepo(sqlDB)
)

type UserService interface {
	Register(u *user.User) (*int64, error)
	Login(u *user.User) error
	// Logout()
}

type userService struct {
	repo user_repo.UserRepo
}

func NewUserService() UserService {
	return &userService{
		repo: userRepo,
	}
}

func (us userService) Register(u *user.User) (*int64, error) {
	err := us.checkNick(u.Nick)
	if err != nil {
		return nil, err
	}
	u.Nick = strings.ToLower(u.Nick)
	u.Password, err = us.saltPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.RoleID = Anon // every new user is registered with 'anon' role
	return us.repo.Register(u)
}
func (us userService) Login(u *user.User) error {
	passwordFromDB, err := us.repo.GetPassword(u.Nick)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(u.Password))
	if err != nil {
		return errs.ErrWrongPassword
	}
	return nil
}

func (us userService) checkNick(nick string) error {
	s := []rune(nick)
	for i := 0; i < len(s); i++ {
		if ok, _ := regexp.MatchString(AllowedChars, string(s[i])); !ok {
			return errs.ErrInvalidNick
		}
	}
	if len(s) < 4 || len(s) > 25 {
		return errs.ErrNickLength
	}
	return nil
}

func (us userService) saltPassword(password string) (string, error) {
	salted, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return "", nil
	}
	password = string(salted)
	return password, nil
}
