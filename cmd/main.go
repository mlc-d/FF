package main

import (
	"gitlab.com/mlc-d/ff/api/web"
	"log"
	"net/http"
)

func main() {
	s := web.NewServer()
	log.Printf("HTTP Server initialized. Listening... (Port %s)\n", s.GetPort())
	log.Fatalln(http.ListenAndServe(s.GetPort(), s))
}
