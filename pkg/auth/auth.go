package auth

import (
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"gitlab.com/mlc-d/ff/pkg/errs"
)

// buildPartialPayload instantiates a new *payload with the given fields
func buildPartialPayload(userID *int64, userRole *uint8) (*payload, error) {
	if userID == nil || userRole == nil {
		return nil, errs.ErrInvalidPayload
	}
	return &payload{
		UserID:   userID,
		UserRole: userRole,
	}, nil
}

const (
	issuer = "gitlab.com/mlc-d"
)

func BuildToken(userID *int64, userRole *uint8) (string, error) {

	signedToken, err := jwt.Sign(t, jwt.WithKey(jwa.HS256, []byte("hey")))
	if err != nil {
		return "", nil
	}
	fmt.Println(string(signedToken[:]))

	return "", nil
}
