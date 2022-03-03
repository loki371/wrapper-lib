package wrmq 

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func BindQueue(ch *amqp.Channel, q amqp.Queue, exchange string) {
	BindQueueWithRouteKey(ch, q, q.Name, exchange)
}

func BindQueueWithRouteKey(ch *amqp.Channel, q amqp.Queue, rKey string, exchange string) {
	err := ch.QueueBind(
		q.Name,   // queue name
		rKey,     // routing key
		exchange, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
}

func GetQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Fail to get queue")
	return q
}

func ExchangeDeclare(ch *amqp.Channel, exchange string, typeEx string) {
	err := ch.ExchangeDeclare(
		exchange, // name
		typeEx,   // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")
}

func Consume(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {
	msges, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	return msges
}

func PublishChannel(ch *amqp.Channel, exchange string, reqQ amqp.Queue, msg amqp.Publishing) {
	err := ch.Publish(
		exchange,  // exchange
		reqQ.Name, // routing key
		false,     // mandatory
		false,     // immediate
		msg)
	failOnError(err, "Failed to publish a message")
}

func CreatePublishingMsg(corrId string, q amqp.Queue, content []byte) amqp.Publishing {
	return amqp.Publishing{
		DeliveryMode:  0,
		CorrelationId: corrId,
		ReplyTo:       q.Name,
		Body:          content,
	}
}

func FindResMsgByCorrId(msgs <-chan amqp.Delivery, corrId string) []byte {
	for d := range msgs {
		if corrId == d.CorrelationId {
			return d.Body
		}
	}
	return nil
} 