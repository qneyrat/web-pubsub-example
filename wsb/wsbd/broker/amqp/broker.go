package amqp

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"

	"web-pubsub-example/wsb/wsbd/channel"
	"web-pubsub-example/wsb/wsbd/message"
)

type Handler interface {
	Handle(c *channel.Channel)
}

type Broker struct{}

var AMQPUri string
var AMQPQueueName string
var AMQPQueueBinding string
var AMQPExchange string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AMQPUri = os.Getenv("AMQP_URI")
	AMQPQueueName = os.Getenv("AMQP_QUEUE_NAME")
	AMQPQueueBinding = os.Getenv("AMQP_QUEUE_BINDING")
	AMQPExchange = os.Getenv("AMQP_EXCHANGE")
}

func (b *Broker) Handle(c *channel.Channel) {
	conn, err := amqp.Dial(AMQPUri)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(AMQPQueueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = ch.QueueBind(q.Name, AMQPQueueBinding, AMQPExchange, false, nil)
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
