package wrmq 

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// -----------------------------------------------------------------------------------------------------
// rabbitMQInterace

type rabbitMQConn struct {
	conn *amqp.Connection
}

type IRabbitMQConn interface {
	openConnection(string)
	CreateChannel() *amqp.Channel
	CloseConnection()
}

func (rbc *rabbitMQConn) openConnection(rabbitMQLink string) {
	var err error
	rbc.conn, err = amqp.Dial(rabbitMQLink)
	failOnError(err, "Failed to connect to RabbitMQ")
}

func (rbc *rabbitMQConn) CreateChannel() *amqp.Channel {
	ch, err := rbc.conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func (rbc *rabbitMQConn) CloseConnection() {
	rbc.conn.Close()
}

// -----------------------------------------------------------------------------------------------------
// factory

func CreateRabbitMQInterface(rabbMQLink string) IRabbitMQConn {
	var rabbitMQConn rabbitMQConn = rabbitMQConn{}
	rabbitMQConn.openConnection(rabbMQLink)
	return &rabbitMQConn
}