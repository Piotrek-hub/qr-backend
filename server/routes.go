package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Check Origin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade to websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("[Error while upgrading to websocket]: ", err)
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {

		}
	}(ws)
	Reader(ws)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	err := json.NewEncoder(w).Encode(map[string]string{"hello": "world"})
	if err != nil {
		return
	}
}

func SetupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/qr", socketHandler)
	r.HandleFunc("/", homeHandler)

	handler := cors.Default().Handler(r)

	http.Handle("/", handler)
}
