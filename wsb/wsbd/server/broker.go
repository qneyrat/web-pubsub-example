package server

import 	"chat-example/wsb/wsbd/channel"

type Broker interface {
	Handle(c *channel.Channel)
}
