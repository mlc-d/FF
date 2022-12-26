package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.com/mlc-d/ff/web/api"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router *chi.Mux
	port   string
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
	server = new(Server)
	log.Println(server)
	server.router = chi.NewRouter()
	server.router.Use(middleware.Logger)

	server.router.Mount("/api", api.Router())

	server.port = os.Getenv("PORT")
	if server.port == "" {
		server.port = DefaultPort
	}
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) GetPort() string {
	return s.port
}
