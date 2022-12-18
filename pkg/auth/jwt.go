package auth

import (
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type payload struct {
	userID   *int64
	userRole *uint8
}

const (
	audience          = "ff-audience"
	expirationInHours = 48
)

func createToken(p payload) (*jwt.Token, error) {
	now := time.Now()
	t, err := jwt.NewBuilder().
		Issuer("ff").
		IssuedAt(now).
		Audience([]string{audience}).
		Subject("auth-token").
		Claim("user-id", p.userID).
		Claim("user-role", p.userRole).
		Expiration(now.Add(time.Hour * expirationInHours)).
		Build()
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func signToken(t *jwt.Token) ([]byte, error) {
	signedToken, err := jwt.Sign(*t, jwt.WithKey(jwa.RS256, keys.Private))
	if err != nil {
		return nil, err
	}
	return signedToken, nil
}
