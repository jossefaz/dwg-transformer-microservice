package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type pickFile struct {
	Name string
	Path string
}

type rabbitmq struct {
	conn  *amqp.Connection
	chanL *amqp.Channel
	queue amqp.Queue
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
	}

}

func newRabbit(connString string, queueName string) (instance rabbitmq) {
	conn, err := amqp.Dial(connString)
	handleError(err, "Can't connect to AMQP")

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	queue, err := amqpChannel.QueueDeclare(queueName, true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")
	return rabbitmq{
		conn:  conn,
		chanL: amqpChannel,
		queue: queue,
	}
}

func (rmq rabbitmq) sendMessage(m pickFile) {
	body, err := json.Marshal(m)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}
	err = rmq.chanL.Publish("", rmq.queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
