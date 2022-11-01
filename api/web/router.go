package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"net/http"
	"os"
)

type Server struct {
	router *chi.Mux
	port   string
	jwt    *jwtauth.JWTAuth
}

const (
	DefaultPort = `:8080`
)

var (
	server *Server
)

func NewServer() *Server {
	if server != nil {
		return server
	}
	server.router = chi.NewRouter()
	server.router.Use(middleware.Logger)

	server.port = os.Getenv("PORT")
	if server.port == "" {
		server.port = DefaultPort
	}
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func (s *Server) GetPort() string {
	return s.port
}
