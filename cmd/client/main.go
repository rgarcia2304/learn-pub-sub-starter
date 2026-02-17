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
	
	queueName := uname + "." + routing.Pausekey
	pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilDirect, 
		queueName,
		routing.Pausekey, 
		amqp.Transient,
	)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

}
