package ws

import "github.com/gorilla/websocket"

const (
	TextType   int = 1
	BinaryType int = 2
)

type Server struct {
	conns map[*websocket.Conn]bool
}

type EventHandler struct {
	OnMessage    func(conn *websocket.Conn, msgType int, msg []byte)
	OnDisconnect func(conn *websocket.Conn)
}
