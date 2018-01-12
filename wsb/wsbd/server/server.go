package server

import (
	"net/http"

	"github.com/qneyrat/wsb/wsbd/channel"
	"github.com/qneyrat/wsb/wsbd/client"
	"github.com/qneyrat/wsb/wsbd/message"
)

type Server struct {
	Clients  client.Clients
	Channels channel.Channels
	Broker Broker
}

type Broker interface {
	Handle(c channel.Channel)
}

func NewServer(b Broker) *Server {
	c := &channel.Channel{
		ID:      "all",
		Chan:    make(chan message.Message),
		Clients: make(client.Clients),
	}

	s := &Server{
		Clients:  make(client.Clients),
		Channels: make(channel.Channels),
	}

	s.AddChannel(c)
	s.Broker = b

	return s
}

func (s *Server) AddChannel(c *channel.Channel) {
	s.Channels[c.ID] = c
}

func (s *Server) Start() error {
	go s.Broker.Handle(s.Channels["all"])
	go s.handleMessages()

	http.HandleFunc("/websocket", s.handleConnections)

	return http.ListenAndServe(":4000", nil)
}
