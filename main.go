package main

import (
	"flag"
	"log"
	"net/http"
	server "qr-backend/server"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	server.SetupRoutes()

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(*addr, nil))
}
