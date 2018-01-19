package message

type Message struct {
	From string
	To   string
	Body []byte
}

func NewMessage(from string, to string, body []byte) Message {
	return Message{
		From: from,
		To:   to,
		Body: body,
	}
}
