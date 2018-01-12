package main

import (
	"log"

	"github.com/qneyrat/wsb/wsbd/server"
)

func main() {
	wbd := server.NewServer()
	err := wbd.Start()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
