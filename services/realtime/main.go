package main

import (
	"fmt"
	"log"
	"net/http"

	"vou/pkg/ws"

	"github.com/gorilla/websocket"
)

func main() {
	log.Println("--> hello from realtime service")

	server := ws.NewServer()

	handler := ws.EventHandler{
		OnMessage: func(conn *websocket.Conn, msgType int, msg []byte) {
			server.Broadcast(msgType, []byte(fmt.Sprintf("%s: %s", conn.RemoteAddr().String(), msg)))
		},
		OnDisconnect: func(conn *websocket.Conn) {
			server.Broadcast(ws.TextType, []byte(fmt.Sprintf("%s disconnected", conn.RemoteAddr().String())))
		},
	}

	http.HandleFunc("/ws", server.CreateHandler(handler))

	log.Fatal(http.ListenAndServe("localhost:8085", nil))
}
