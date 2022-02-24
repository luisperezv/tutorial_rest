package main

import (
	"log"
	"net/http"

	"tutorial_rest/pkg/server"
)

func main() {

	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))

}
