package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gitlab.com/mlc-d/ff/pkg/auth"
	"gitlab.com/mlc-d/jam"
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

	_ = auth.NewJWTService()
	keys := auth.GetKeys()

	j, err := jam.New(
		jam.RS256,
		jam.DefaultLookupOptions,
		keys.Private,
		keys.Public,
		jam.TokenFromCookie,
		jam.TokenFromHeader)
	if err != nil {
		panic(err)
	}

	protected := chi.NewRouter()
	protected.Use(jam.Verifier(j))
	protected.Use(jam.Authenticator)
	registerRoutes(protected, protectedRoutes)

	server.router.Mount("/protected", protected)

	server.port = os.Getenv("PORT")
	if server.port == "" {
		server.port = DefaultPort
	}
	registerRoutes(server.router, authenticationRoutes)
	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) GetPort() string {
	return s.port
}
