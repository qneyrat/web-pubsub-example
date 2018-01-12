package message

type Message struct {
	From string
	Body []byte
}

func NewMessage(from string, body []byte) Message {
	return Message{
		From: from,
		Body: body,
	}
}
