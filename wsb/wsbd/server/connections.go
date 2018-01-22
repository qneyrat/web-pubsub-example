package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/qneyrat/wsb/wsbd/client"
)

func (s *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(sessionKey).(Session).Identifier
	log.Printf("New client connected with ID %v!", id)

	upgrader := websocket.Upgrader{
		EnableCompression: true,
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	s.Clients[id] = &client.Client{
		ID:   id,
		Conn: ws,
	}

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
	}
}
