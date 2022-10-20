package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/mlc-d/ff/api/web"
)

const (
	DefaultPort = `:8080`
)

func main() {
	r := web.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
	log.Printf("Listening... (Port %s)\n", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
