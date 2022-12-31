package main

import (
	"gitlab.com/mlc-d/ff/web/router"
	"log"
	"net/http"
)

func main() {
	s := router.NewServer()
	log.Printf("HTTP Server initialized. Listening... (Port %s)\n", s.GetPort())
	log.Fatalln(http.ListenAndServe(s.GetPort(), s))
}
