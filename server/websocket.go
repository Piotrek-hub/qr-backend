package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"qr-backend/utils"
)

type request struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

var hubs = make(map[string]*websocket.Conn)

func Reader(ws *websocket.Conn) {
	for {
		messageType, msg, err := ws.ReadMessage()

		if err != nil {
			log.Println("[Error during reading message]: ", err)
			break
		}

		// Load request to struct
		var req request
		err = json.Unmarshal(msg, &req)
		if err != nil {
			log.Println("err: ", err)
			break
		}

		//Handle request type
		switch req.Message {
		case "requestToken":
			token := utils.GenerateToken()
			resp, _ := json.Marshal(map[string]string{"token": token})

			err := ws.WriteMessage(messageType, resp)
			if err != nil {
				log.Println("[Error during writing message]", err)
			}

			hubs[token] = ws
		case "checkToken":
			resp, _ := json.Marshal(map[string]string{"message": "unlock"})
			err := hubs[req.Token].WriteMessage(messageType, resp)
			if err != nil {
				log.Println(err)
			}

			err = ws.WriteMessage(messageType, []byte("Device unlocked"))
			if err != nil {
				log.Println(err)
			}
		default:
			log.Println("Unknown message type")
		}

	}
}
