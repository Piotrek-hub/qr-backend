package utils

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type request struct {
	Message string
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("[Error during reading message]: ", err)
			return
		}

		// Load request to struct
		req := request{}
		err = json.Unmarshal(message, &req)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Handle request type
		switch req.Message {
		case "requestToken":
			token := GenerateToken()
			resp, _ := json.Marshal(map[string]string{"token": token})

			err := conn.WriteMessage(messageType, resp)
			if err != nil {
				log.Fatal("[Error during writing message]", err)
				return
			}

		default:
			log.Fatal("Unknown message type")
		}

	}
}
