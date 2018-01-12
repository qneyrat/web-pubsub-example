package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/qneyrat/wsb/wsbd/client"
)

func (s *Server) handleConnections(w http.ResponseWriter, r *http.Request) {
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

	//	id := uuid.NewV4().String()
	id := "all"
	s.Clients[id] = &client.Client{
		ID:   id,
		Conn: ws,
	}

	log.Printf("new client connected with ID %v", id)

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
	}
}
