package pubsub

import(
	amqp "github.com/rabbitmq/amqp091-go"
	"encoding/json"
	"errors"
	"log"
)

type AckType int
const(
	Ack AckType = iota
	NackRequeue
	NackDiscard

)

func SubscribeJSON[T any](
		conn *amqp.Connection, 
		exchange, 
		queueName, 
		key string, 
		queueType SimpleQueueType,
		handler func(T) AckType,
) error{
	c, _, err := DeclareAndBind(conn, exchange, queueName, key, queueType)

	if err != nil{
		return errors.New("Error after declaring and binding")
	}

	deliveries, err := c.Consume(queueName,"", false, false, false, false, nil)

	if err != nil{
		log.Println(err)
	}

	go func(){
		for dat := range deliveries{
			var output T
			if err := json.Unmarshal(dat.Body, &output) ; err != nil{
				log.Println("Message is not acknowledged")
				continue
			}
			ackState := handler(output)
			
			switch ackState {
				case Ack:
					_ = dat.Ack(false)
					log.Println("Message has been Ack'd")
				case NackRequeue:
					_  = dat.Nack(false, true)
					log.Println("Message has not been Ack'd will retry")
				case NackDiscard: 
					_ = dat.Nack(false, false)
					log.Println("Message has not been Ack'd will discard")
			}
		}
	}()
	
	return nil
}

	

