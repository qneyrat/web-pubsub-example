package main

import (
	"log"
	"github.com/qneyrat/wsb/wsbd/broker"
	"github.com/qneyrat/wsb/wsbd/server"
)

func main() {
	wbd := server.NewServer(&broker.HttpBroker{})
	err := wbd.Start()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
