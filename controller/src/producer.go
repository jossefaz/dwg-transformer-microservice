package main

import (
	"encoding/json"
	"fmt"
	"os"

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
	conn := dial(connString)
	amqpChannel := getChannel(conn)
	queue := connectToQueue(amqpChannel, queueName)
	return rabbitmq{
		conn:  conn,
		chanL: amqpChannel,
		queue: queue,
	}
}

func dial(connString string) *amqp.Connection {
	conn, err := amqp.Dial(connString)
	handleError(err, "Can't connect to AMQP")
	if err != nil {
		os.Exit(1)
	}
	return conn
}

func getChannel(conn *amqp.Connection) *amqp.Channel {
	c, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")
	if err != nil {
		os.Exit(1)
	}
	return c
}

func connectToQueue(c *amqp.Channel, queueName string) amqp.Queue {
	q, err := c.QueueDeclare(queueName, true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")
	if err != nil {
		os.Exit(1)
	}
	return q
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
