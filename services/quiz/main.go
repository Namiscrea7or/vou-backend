package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"vou/pkg/ws"
)

func main() {
	log.Println("--> hello from realtime service")

	server := ws.NewServer()

	handler := ws.EventHandler{
		OnMessage: func(ctx context.Context, msgType int, msg []byte) {
			if msgType == ws.BinaryType {
				return
			}

			var parsedMessage ws.Message
			err := json.Unmarshal(msg, &parsedMessage)
			if err != nil {
				log.Println("--> failed to parse message", err)
				return
			}

			log.Println("--> parsed message", parsedMessage)
		},
		OnDisconnect: func(ctx context.Context) {
			idToken := ctx.Value(ws.AuthKey)
			server.Broadcast(ws.TextType, []byte(fmt.Sprintf("%s disconnected", idToken)))
		},
	}

	http.HandleFunc("/ws", server.CreateHandler(handler))

	log.Fatal(http.ListenAndServe("localhost:8085", nil))
}
