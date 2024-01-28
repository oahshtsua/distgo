package main

import (
	"log"

	"github.com/oahshtsua/distgo/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":5000")
	log.Fatal(srv.ListenAndServe())
}
