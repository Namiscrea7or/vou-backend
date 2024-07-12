package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	TextType   int = 1
	BinaryType int = 2
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleConnect(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("failed to upgrade protocol:", err)
		return
	}
	defer conn.Close()

	s.conns[conn] = true
	s.broadcast(TextType, []byte(fmt.Sprintf("%s connected", conn.RemoteAddr().String())))

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			delete(s.conns, conn)
			s.broadcast(TextType, []byte(fmt.Sprintf("%s disconnected", conn.RemoteAddr().String())))
			log.Printf("failed to read msg from %s: %s\n", conn.RemoteAddr(), err)
			break
		}

		s.broadcast(msgType, []byte(fmt.Sprintf("%s: %s", conn.RemoteAddr().String(), msg)))
	}
}

func (s *Server) broadcast(msgType int, msg []byte) {
	for conn := range s.conns {
		err := conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Printf("failed to send msg to %s: %s\n", conn.RemoteAddr(), err)
			continue
		}
	}
}

func main() {
	log.Println("--> hello from realtime service")

	server := NewServer()
	http.HandleFunc("/ws", server.handleConnect)
	log.Fatal(http.ListenAndServe("localhost:8085", nil))
}
