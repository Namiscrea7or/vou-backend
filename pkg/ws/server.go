package ws

import (
	"context"
	"log"
	"net/http"
	"strings"

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

		idToken := ""
		authKey := r.Header.Get("Authorization")
		if strings.HasPrefix(authKey, "Bearer ") {
			idToken = strings.Replace(authKey, "Bearer ", "", 1)
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, AuthKey, idToken)
		ctx = context.WithValue(ctx, ConnKey, conn)

		s.conns[conn] = true
		log.Printf("%s connected: %s", conn.RemoteAddr().String(), idToken)

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				delete(s.conns, conn)
				log.Printf("failed to read msg from %s: %s\n", conn.RemoteAddr(), err)
				go handler.OnDisconnect(ctx)
				break
			}

			go handler.OnMessage(ctx, msgType, msg)
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

func (s *Server) Respond(conn *websocket.Conn, msgType int, msg []byte) error {
	err := conn.WriteMessage(msgType, msg)
	if err != nil {
		return err
	}

	return nil
}
