package auth

import "github.com/lestrrat-go/jwx/v2/jwt"

type JWTRepo interface {
	start()
	// stop()
	create(p payload) ([]byte, error)
	verify(t *jwt.Token) error
}

type jwtRepo struct {
	keys *jwkSet
}

func new() JWTRepo {
	var jr jwtRepo
	jr.start()
	return &jr
}

func (jr *jwtRepo) start() {
	jr.keys = getKeys()
}

// func (jr *jwtRepo) stop() {
// 	jr.keys.private = nil
// 	jr.keys.public = nil
// }

func (jr *jwtRepo) create(p payload) ([]byte, error) {
	t, err := createToken(p)
	if err != nil {
		return nil, err
	}
	st, err := signToken(t)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (jr *jwtRepo) verify(t *jwt.Token) error {
	return nil
}
