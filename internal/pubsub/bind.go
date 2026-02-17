package pubsub

import(
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func DeclareAndBind(
	conn *amqp.Connection, 
	exchange, 
	queueName, 
	key string, 
	queueType SimpleQueueType, 
) (*ampq.Channel, ampq.Queue, error){

	channel, err := conn.Channel()
	if err != nil{
		log.Fatalf("Failed to bind the queue" ,err)
	}
	
	var isDurable bool
	if queueType == ampq.Durable{
		isDurable = true
	}else{
		isDurable = false
	}

	var isTransient
	var isExclusive
	if queueType == ampq.Transient{
		isTransient = true
		isExclusive = true
	}else{
		isTransient = false
		isExclusive = false
	}


	queue, err := channel.QueueDeclare(
		queueName, 
		isDurable,
		isTransient,
		isExclusive, 
		false,
		nil,
	)

	if err != nil{
		log.Fatalf("Failed to declare queue", err)
	}

	if err = channel.QueueBind(queueName, key, exchange, false, nil)
	if err != nil{
		log.Fatalf("Failed to bind the queue", err)
	}

	return channel, queue, nil 
}
