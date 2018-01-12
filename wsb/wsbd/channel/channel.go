package channel

import (
	"github.com/qneyrat/wsb/wsbd/client"
	"github.com/qneyrat/wsb/wsbd/message"
)

type Channel struct {
	ID      string
	Chan    chan message.Message
	Clients client.Clients
}

type Channels map[string]*Channel
