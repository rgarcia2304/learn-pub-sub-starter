package main

import (
	"log"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/pubsub"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/routing"
	"github.com/rgarcia2304/learn-pub-sub-starter/internal/gamelogic"

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
	
	if err != nil{
		log.Fatalf("Program failed because of %v", err)
	}

	queueName := routing.GameLogSlug

	_, _ , err = pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilTopic, 
		queueName,
		"game_logs.*", 
		pubsub.Durable,
	)

	for {
		words := gamelogic.GetInput()

		if len(words) == 0{
			continue
		}

		switch words[0]{
			case "pause":
				log.Printf("Sending a pause message")
				pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})
			case "resume":
				log.Printf("Sending a resume message")
				pubsub.PublishJSON(channel, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: false})
			case "quit":
				log.Printf("Exiting ......")
				return
			default:
				log.Printf("Dont understand the command")

		}
	}
	

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	
}
