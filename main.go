package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	setupRoutes()

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(*addr, nil))
}
