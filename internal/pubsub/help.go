package pubsub

import(
	amqp "github.com/rabbitmq/amqp091-go"
	"encoding/json"
	"context"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error{
	
	//marshall the val to json bytes 
	marshalledVal, err := json.Marshal(val)
	if err != nil{
		return err
	}
	msg := amqp.Publishing{ContentType: "application/json", Body: marshalledVal}
	_ =  ch.PublishWithContext(context.Background(), exchange, key, false, false, msg)
	return nil

	

}
