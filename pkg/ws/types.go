package ws

import (
	"context"

	"github.com/gorilla/websocket"
)

const (
	TextType   int = 1
	BinaryType int = 2
)

type Server struct {
	conns map[*websocket.Conn]bool
}

type EventHandler struct {
	OnMessage    func(ctx context.Context, msgType int, msg []byte)
	OnDisconnect func(ctx context.Context)
}

type MessageType string

type Message struct {
	Event   string `json:"event"`
	Payload any    `json:"payload"`
}

type ContextKey int

const (
	AuthKey ContextKey = iota
	ConnKey
)
