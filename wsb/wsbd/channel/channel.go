package channel

import (
	"github.com/qneyrat/wsb/wsbd/message"
)

type Channel struct {
	Chan chan message.Message
}
