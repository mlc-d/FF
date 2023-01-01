package user

import (
	"gitlab.com/mlc-d/ff/dto"
	"gitlab.com/mlc-d/ff/pkg/user/internal"
	"regexp"
	"strings"

	"gitlab.com/mlc-d/ff/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

const (
	Dev = iota
	Admin
	Mod
	Anon

	AllowedChars = `[\p{L}\p{N}]` // Only alphanumeric characters are allowed
)

var (
	userRepo = internal.NewRepo()
)

type Service interface {
	Register(u *dto.User) (*int64, *uint8, error)
	Login(u *dto.User) (*int64, *uint8, error)
	checkPasswordByNick(nick, password string) (*int64, *uint8, error)
}

type service struct {
	repo internal.Repo
}

func NewService() Service {
	return &service{
		repo: userRepo,
	}
}

func (us service) Register(u *dto.User) (*int64, *uint8, error) {

	user := new(internal.User)
	user.Nick = u.Nick
	user.Password = u.Password

	err := us.checkNick(user.Nick)
	if err != nil {
		return nil, nil, err
	}
	user.Nick = strings.ToLower(user.Nick)
	user.Password, err = us.saltPassword(user.Password)
	if err != nil {
		return nil, nil, err
	}
	user.RoleID = Anon // every new user is registered with 'anon' role
	return us.repo.Register(user)
}
func (us service) Login(u *dto.User) (*int64, *uint8, error) {
	return us.checkPasswordByNick(strings.ToLower(u.Nick), u.Password)
}

func (us service) checkPasswordByNick(nick, password string) (*int64, *uint8, error) {
	id, role, passwordFromDB, err := us.repo.GetPassword(nick)
	if err != nil {
		return nil, nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(password))
	if err != nil {
		return nil, nil, errs.ErrInvalidCredentials
	}
	return id, role, nil
}

func (us service) checkNick(nick string) error {
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

func (us service) saltPassword(password string) (string, error) {
	salted, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return "", nil
	}
	password = string(salted)
	return password, nil
}
