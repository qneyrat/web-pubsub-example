package message

type Message struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

func NewMessage(from string, to string, body string) Message {
	return Message{
		From: from,
		To:   to,
		Body: body,
	}
}
