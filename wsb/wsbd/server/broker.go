package server

import "web-pubsub-example/wsb/wsbd/channel"

type Broker interface {
	Handle(c *channel.Channel)
}
