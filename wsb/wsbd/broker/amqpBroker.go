package broker

import (
	"log"

	"github.com/streadway/amqp"

	"github.com/qneyrat/wsb/wsbd/channel"
	"github.com/qneyrat/wsb/wsbd/message"
)

type AmqpBroker struct{}

func (b *AmqpBroker) Handle(c *channel.Channel) {
	conn, err := amqp.Dial("amqp://admin:rabbitmq@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"messages", // name of the queue
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = ch.QueueBind(
		q.Name, // queue name
		"api.conversation.*.message.added", // routing key
		"default",                          // exchange
		false,
		nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("%s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
			message := message.Message{
				From: "all",
				Body: d.Body,
			}

			c.Chan <- message
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
