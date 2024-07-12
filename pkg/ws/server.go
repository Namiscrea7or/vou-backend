package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// For development purpose, we allow all connections
		return true
	},
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) CreateHandler(handler EventHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("failed to upgrade protocol:", err)
			return
		}
		defer conn.Close()

		s.conns[conn] = true
		s.Broadcast(TextType, []byte(fmt.Sprintf("%s connected", conn.RemoteAddr().String())))

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				delete(s.conns, conn)
				log.Printf("failed to read msg from %s: %s\n", conn.RemoteAddr(), err)
				go handler.OnDisconnect(conn)
				break
			}

			go handler.OnMessage(conn, msgType, msg)
		}
	}
}

func (s *Server) Broadcast(msgType int, msg []byte) {
	for conn := range s.conns {
		err := conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Printf("failed to send msg to %s: %s\n", conn.RemoteAddr(), err)
			continue
		}
	}
}
