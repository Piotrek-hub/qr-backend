package main

import (
	"crypto/rand"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"math/big"
	"net/http"
)

type request struct {
	Message string
}

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Origin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// Upgrade to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("[Error while upgrading to websocket]: ", err)
		return
	}
	defer conn.Close()
	log.Println("New Connection")
	// Read messages
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("[Error during reading message]: ", err)
			return
		}
		req := request{}
		err = json.Unmarshal(message, &req)

		switch req.Message {
		case "requestToken":
			token, _ := rand.Prime(rand.Reader, 64)
			resp, _ := json.Marshal(map[string]*big.Int{"token": token})
			err := conn.WriteMessage(messageType, []byte(resp))
			if err != nil {
				log.Println("[Error during writing message]", err)
				return
			}

		default:
			log.Println("Unknown message type")
		}

		log.Printf("message: %s\n", req)
		log.Printf("message type: %d\n", messageType)
	}
}

func setupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/qr", socketHandler)

	handler := cors.Default().Handler(r)

	http.Handle("/", handler)
}
