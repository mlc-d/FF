package internal

import (
	keys2 "gitlab.com/mlc-d/ff/pkg/keys"
	user_service "gitlab.com/mlc-d/ff/pkg/user"
	"gitlab.com/mlc-d/go-jam"
)

var (
	UserService   = user_service.NewUserService()
	keys          = keys2.GetKeys()
	JWTService, _ = jam.New(
		jam.RS256,
		jam.DefaultLookupOptions,
		keys.Private,
		keys.Public,
		jam.TokenFromCookie,
		jam.TokenFromHeader,
	)
)
