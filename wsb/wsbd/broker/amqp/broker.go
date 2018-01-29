package amqp

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"

	"chat-example/wsb/wsbd/channel"
	"chat-example/wsb/wsbd/message"
)

type Handler interface {
	Handle(c *channel.Channel)
}

type Broker struct{}

func (b *Broker) Handle(c *channel.Channel) {
	conn, err := amqp.Dial("amqp://admin:admin@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("messages", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = ch.QueueBind(q.Name, "api.conversation.*.message.*.added", "api", false, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)

			data := &message.Message{}
			err := json.Unmarshal(d.Body, data)
			if err != nil {
				log.Fatalf("%s", err)
			}

			message := message.Message{
				From: data.From,
				To:   data.To,
				Body: string(d.Body),
			}

			c.Chan <- message
		}
	}()
	<-forever
}
