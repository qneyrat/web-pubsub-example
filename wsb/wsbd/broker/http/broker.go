package http

import (
	"chat-example/wsb/wsbd/channel"
	"chat-example/wsb/wsbd/message"
	"encoding/json"
	"log"
	"net/http"
)

type Handler interface {
	Handle(c *channel.Channel)
}

type Broker struct{}

func (b *Broker) Handle(c *channel.Channel) {
	http.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		str := `{"from": "test2", "to": "test", "body": "1"}`

		data := &message.Message{}
		err := json.Unmarshal([]byte(str), data)
		if err != nil {
			log.Fatalf("%s", err)
		}

		message := message.Message{
			From: data.From,
			To:   data.To,
			Body: str,
		}

		c.Chan <- message
	})

	log.Fatal(http.ListenAndServe(":8089", nil))
}
