package web

import (
	"github.com/go-chi/chi/v5"
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
	authenticationRoutes routes = []*route{{
		method:  "POST",
		pattern: "/register",
		// handler: user_service,
	}}
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
