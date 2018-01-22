package broker

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/qneyrat/wsb/wsbd/channel"
	"github.com/qneyrat/wsb/wsbd/message"
)

type HttpBroker struct{}

func (b *HttpBroker) Handle(c *channel.Channel) {
	http.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		str := []byte(`{"from": "1", "to": "test", "body": "1"}`)

		data := &ApiMessage{}
		err := json.Unmarshal(str, data)
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
