package web

import (
	"github.com/go-chi/chi/v5"
	"gitlab.com/mlc-d/ff/api/web/handler"
	"log"
	"net/http"
)

type route struct {
	method  string
	pattern string
	handler http.HandlerFunc
}

type routes []*route

var (
	authenticationRoutes routes = []*route{
		{
			method:  "POST",
			pattern: "/register",
			handler: handler.RegisterUser,
		},
		{
			method:  "GET",
			pattern: "/api/users",
			handler: func(writer http.ResponseWriter, request *http.Request) {
				_, _ = writer.Write([]byte(`test`))
			},
		},
		{
			method:  "POST",
			pattern: "/login",
			handler: handler.Login,
		},
	}
)

func registerRoutes(mux *chi.Mux, routes routes) {
	for _, r := range routes {
		switch r.method {
		case "GET":
			mux.Get(r.pattern, r.handler)
		case "POST":
			mux.Post(r.pattern, r.handler)
		case "PATCH":
			mux.Patch(r.pattern, r.handler)
		case "DELETE":
			mux.Delete(r.pattern, r.handler)
		default:
			log.Printf("Cannot register route: %s (invalid method %s)\n",
				r.pattern,
				r.method)
		}
	}
}
