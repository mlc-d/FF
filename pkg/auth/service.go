package auth

import "github.com/lestrrat-go/jwx/v2/jwt"

type KeyService interface {
	CreateToken(userID *int64, userRole *uint8) ([]byte, error)
	VerifyToken(t *jwt.Token) error
}

var repo = new()

type keyService struct {
	repo JWTRepo
}

// func (ks *keyService) start() {
// 	ks.repo.start()
// }

// func (ks *keyService) stop() {
// 	ks.repo.stop()
// }

func (ks *keyService) CreateToken(userID *int64, userRole *uint8) ([]byte, error) {
	p := payload{
		userID:   userID,
		userRole: userRole,
	}
	return ks.repo.create(p)
}

func (ks *keyService) VerifyToken(t *jwt.Token) error {
	return nil
}

func NewKeyService() KeyService {
	var ks keyService
	ks.repo = repo
	// ks.start()
	return &ks
}
