package main

import (
	"log"

	"chat-example/wsb/wsbd/broker"
	"chat-example/wsb/wsbd/server"
)

func main() {
	wbd := server.NewServer(&broker.AMQPBroker{})
	err := wbd.Start()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
