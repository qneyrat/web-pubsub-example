package channel

import (
	"chat-example/wsb/wsbd/message"
)

type Channel struct {
	Chan chan message.Message
}
