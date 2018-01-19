package server

import "log"

func (s *Server) handleMessages() {
	for {
		message := <-s.Channel.Chan
		log.Printf(" [v] %s", message.Body)
		//client := s.Clients[message.From]
		//ws := client.Conn
		//err := ws.WriteMessage(1, message.Body)
		//if err != nil {
		//	log.Printf("error: %v", err)
		//	ws.Close()
		//	delete(s.Clients, client.ID)
		//}
	}
}
