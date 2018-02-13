package channel

import (
	"web-pubsub-example/wsb/wsbd/message"
)

type Channel struct {
	Chan chan message.Message
}
