package main

import (
	"log"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/pubsub"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/routing"
)

func main() {
	fmt.Println("Starting Peril server...")
	connectionString := "amqp://guest:guest@127.0.0.1:5672/"
	conn, err := amqp.Dial(connectionString)
	if err != nil{
		log.Fatalf("Program failed because of %v", err)
	}
	defer conn.Close()

	fmt.Println("Connection Was Successful")

	channel, err := conn.Channel()
	
	pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})

	if err != nil{
		log.Fatalf("Program failed because of %v", err)
	}


	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	
}
