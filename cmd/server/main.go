package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/pubsub"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/routing"
)

func main() {
	fmt.Println("Starting Peril server...")
	connectionString := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connectionString)
	if err != nil{
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("Connection Was Successful")

	channel, err := connection.Channel()
	
	pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})

	if err != nil{
		fmt.Println(err)
	}


	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	
}
