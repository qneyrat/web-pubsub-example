package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	
	"github.com/joho/godotenv"

	"chat-example/wsb/wsbd/broker/http"
	"chat-example/wsb/wsbd/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	pubKeyPath := os.Getenv("JWT_KEY")
	log.Println(pubKeyPath)


	errs := make(chan error, 2)
	go func() {
		log.Println("wsbd start")
		errs <- server.NewServer(&http.Broker{}).Start()
	}()
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		errs <- fmt.Errorf("%s", <-stop)
	}()

	log.Println("terminated", <-errs)
}
