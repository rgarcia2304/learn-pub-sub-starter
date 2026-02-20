package pubsub

import(
	amqp "github.com/rabbitmq/amqp091-go"
	"fmt"
)

type SimpleQueueType int

const(
	Transient SimpleQueueType = iota
	Durable 
)

func DeclareAndBind(
	conn *amqp.Connection, 
	exchange, 
	queueName, 
	key string, 
	queueType SimpleQueueType, 
) (*amqp.Channel, amqp.Queue, error){

	channel, err := conn.Channel()
	if err != nil{
		return nil, amqp.Queue{}, fmt.Errorf("open channel: %w", err)
	}
	
	
	durable := queueType == Durable
	autoDelete := queueType == Transient
	exclusive := queueType == Transient

	queue, err := channel.QueueDeclare(
		queueName, 
		durable,
		autoDelete,
		exclusive, 
		false,
		nil,
	)

	if err != nil{
		channel.Close()
		return nil, amqp.Queue{}, fmt.Errorf("declare queue %w ", err)
	}

	if err := channel.QueueBind(queueName, key, exchange, false, nil); err != nil{
		channel.Close()
		return nil, amqp.Queue{}, fmt.Errorf("bind queue : %w", err)
	}

	return channel, queue, nil 
}
