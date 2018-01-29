package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"chat-example/wsb/wsbd/broker/amqp"
	"chat-example/wsb/wsbd/server"
)

func main() {
	errs := make(chan error, 2)
	go func() {
		log.Println("wsbd start")
		errs <- server.NewServer(&amqp.Broker{}).Start()
	}()
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		errs <- fmt.Errorf("%s", <-stop)
	}()

	log.Println("terminated", <-errs)
}
