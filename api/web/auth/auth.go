package auth

import (
	"log"

	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func NewTokenAuth() *jwtauth.JWTAuth {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	log.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
	return tokenAuth
}
