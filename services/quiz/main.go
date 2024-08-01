package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"vou/pkg/quiz"
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

			response, err := quiz.HandleGameEvent(ctx, parsedMessage)
			if err != nil {
				response = &quiz.Response{
					Status: quiz.Failure,
					Event:  quiz.Event(parsedMessage.Event),
					Payload: map[string]interface{}{
						"message": err.Error(),
					},
				}
			}

			byteResponse, err := json.Marshal(response)
			if err != nil {
				log.Println("--> failed to marshal response", response, err)
				return
			}

			// NOTE: instead of broadcast to all clients, we should implement
			// observer pattern based on what clients care for
			// e.g. A game session's events should only be sent to clients who are in that game
			server.Broadcast(ws.TextType, byteResponse)
		},
		OnDisconnect: func(ctx context.Context) {
			idToken := ctx.Value(ws.AuthKey)
			server.Broadcast(ws.TextType, []byte(fmt.Sprintf("%s disconnected", idToken)))
		},
	}

	http.HandleFunc("/ws", server.CreateHandler(handler))

	log.Fatal(http.ListenAndServe("localhost:8085", nil))
}
