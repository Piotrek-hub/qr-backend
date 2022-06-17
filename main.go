package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "htp service address")

func main() {
	setupRoutes()

	log.Println("Listening on localhost:8080")

	log.Println(http.ListenAndServe(*addr, nil))
}
