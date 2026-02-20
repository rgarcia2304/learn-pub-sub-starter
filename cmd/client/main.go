package main

import(
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
	fmt.Println("Starting Peril client...")
	connectionString := "amqp://guest:guest@127.0.0.1:5672/"

	conn, err := amqp.Dial(connectionString)
	if err != nil{
		log.Fatalf("Program failed because of %v", err)
	}
	defer conn.Close()

	uname, err := gamelogic.ClientWelcome()

	
	if err != nil{
		log.Fatalf("Issue getting name %v", err)
	}
	
	channel, err := conn.Channel()
	//create a new gamestate
	newGameState := gamelogic.NewGameState(uname) 
	pauseName := routing.PauseKey + "." + uname
	armyKey := routing.ArmyMovesPrefix + "." + "*"
	armyQueue := routing.ArmyMovesPrefix + "." + uname

	err = pubsub.SubscribeJSON(
		conn, 
		routing.ExchangePerilDirect, 
		pauseName, 
		routing.PauseKey, 
		pubsub.Transient, 
		handlerPause(newGameState),
	)

	err = pubsub.SubscribeJSON(
		conn, 
		routing.ExchangePerilTopic, 
		armyQueue, 
		armyKey, 
		pubsub.Transient, 
		func(mv gamelogic.ArmyMove) {
        	newGameState.HandleMove(mv)
        	fmt.Print("> ")
    		},
	)

	if err != nil{
		log.Fatalf("Issue declaring the queue")
	}


	for{
		words := gamelogic.GetInput()

		if len(words) == 0{
			continue
		}

		switch words[0]{
			case "spawn":
				newGameState.CommandSpawn(words)
			case "move":
				mv, err := newGameState.CommandMove(words)
				if err != nil{
					log.Println("issue making move")
					break
				}
				pubsub.PublishJSON(channel, routing.ExchangePerilTopic, armyQueue , mv)
				log.Println("Message Was successfully printed")


			case "status":
				newGameState.CommandStatus()
			case "help":
				gamelogic.PrintClientHelp()
			case "spam":
				log.Println("Spamming is not allowed yet")
			case "quit":
				gamelogic.PrintQuit()
				return
			default:
				log.Println("Error: Command not known see help for list of available commands")
		}
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

}
