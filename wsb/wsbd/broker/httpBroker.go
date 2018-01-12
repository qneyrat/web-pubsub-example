package broker

import (
	"log"
	"net/http"

	"github.com/qneyrat/wsb/wsbd/channel"
	"github.com/qneyrat/wsb/wsbd/message"
)

type HttpBroker struct {}

func (b* HttpBroker) Handle(c channel.Channel) {
	http.HandleFunc("/actions", func(w http.ResponseWriter, r *http.Request) {
		str := `{"message": "1"}`
		log.Printf("new message  %v", str)

		body := []byte(str)
		message := message.Message{
			From: "all",
			Body: body,
		}

		c.Chan <- message
	})

	log.Fatal(http.ListenAndServe(":8007", nil))
}
