package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"qr-backend/utils"
)

type request struct {
	Message string
}

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
		log.Println("req: ", req)
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
				log.Fatal("[Error during writing message]", err)
			}

		default:
			log.Fatal("Unknown message type")
		}

	}
}
