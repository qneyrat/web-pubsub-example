package server

import (
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
	handle(c *channel.Channel)
}

func NewServer() *Server {
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

	return s
}

func (s *Server) AddChannel(c *channel.Channel) {
	s.Channels[c.ID] = c
}

func (s *Server) AddBroker(b Broker) {
	s.Broker = b
}

func (s *Server) Start() {
	go s.Broker.handle(s.Channels["all"])
	go s.handleMessages()
}
