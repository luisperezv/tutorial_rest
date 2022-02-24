package main

import (
	"log"
	"net/http"
	//"github.com/friendsofgo/gopherapi/pkg/server"
)

func main() {

	//s := server.New()
	log.Fatal(http.ListenAndServe(":8080", nil))
	//log.Fatal(http.ListenAndServe(":8080", s.Router()))

	// Instanciar tracer global Jaeger
	//traceCfg, err := jaegercfg.conf config.FromEnv()

	//_ = traceCfg
	//_ = err

}
