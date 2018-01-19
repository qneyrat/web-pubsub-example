package server

import (
	"net/http"

	"github.com/qneyrat/wsb/wsbd/channel"
	"github.com/qneyrat/wsb/wsbd/client"
	"github.com/qneyrat/wsb/wsbd/message"
)

type Server struct {
	Clients client.Clients
	Channel *channel.Channel
	Broker  Broker
}

type Broker interface {
	Handle(c *channel.Channel)
}

func NewServer(b Broker) *Server {
	c := &channel.Channel{
		Chan: make(chan message.Message),
	}

	s := &Server{
		Clients: make(client.Clients),
		Channel: c,
		Broker:  b,
	}

	return s
}

func (s *Server) Start() error {
	go s.Broker.Handle(s.Channel)
	go s.handleMessages()

	http.Handle(
		"/websocket",
		jwtMiddleware(http.HandlerFunc(s.handleConnections)),
	)

	return http.ListenAndServe(":4000", nil)
}
