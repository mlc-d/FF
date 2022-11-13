package auth

import (
	"bytes"
	"errors"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jws"
)

type KeyService interface {
	CreateToken(userID *int64, userRole *uint8) ([]byte, error)
	VerifyToken(t []byte) error
}

var repo = newJWTRepo()

type keyService struct {
	repo JWTRepo
}

func (ks *keyService) CreateToken(userID *int64, userRole *uint8) ([]byte, error) {
	p := payload{
		userID:   userID,
		userRole: userRole,
	}
	return ks.repo.create(p)
}

func (ks *keyService) VerifyToken(t []byte) error {
	v, err := jws.Verify(t, jws.WithKey(jwa.RS256, keys.public))
	if err != nil {
		return err
	}
	if !bytes.Equal(t, v) {
		return errors.New("invalid token")
	}
	return nil
}

func NewJWTService() KeyService {
	var ks keyService
	ks.repo = repo
	return &ks
}
