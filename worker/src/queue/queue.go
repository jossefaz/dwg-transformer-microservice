package queue

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"os/exec"
	"src/worker/utils"
)

type rabbitmq struct {
	Conn  *amqp.Connection
	ChanL *amqp.Channel
	Queue amqp.Queue
}



func NewRabbit(connString string, queueName string) (instance rabbitmq) {
	conn := dial(connString)
	amqpChannel := getChannel(conn)
	queue := connectToQueue(amqpChannel, queueName)
	return rabbitmq{
		Conn:  conn,
		ChanL: amqpChannel,
		Queue: queue,
	}
}

func dial(connString string) *amqp.Connection {
	conn, err := amqp.Dial(connString)
	utils.HandleError(err, "Can't connect to AMQP")
	if err != nil {
		os.Exit(1)
	}
	return conn
}

func getChannel(conn *amqp.Connection) *amqp.Channel {
	c, err := conn.Channel()
	utils.HandleError(err, "Can't create a amqpChannel")
	if err != nil {
		os.Exit(1)
	}
	return c
}

func connectToQueue(c *amqp.Channel, queueName string) amqp.Queue {
	q, err := c.QueueDeclare(queueName, true, false, false, false, nil)
	utils.HandleError(err, "Could not declare `add` queue")
	if err != nil {
		os.Exit(1)
	}
	return q
}

func (rmq rabbitmq) sendMessage(body []byte) {
	err := rmq.ChanL.Publish("", rmq.Queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
}

func (rmq rabbitmq) ListenMessage() {

	err := rmq.ChanL.Qos(1, 0, false)
	utils.HandleError(err, "Could not configure QoS")
	messageChannel, err := rmq.ChanL.Consume(
		rmq.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.HandleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		counter := 0
		fmt.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			counter++
			fmt.Printf("\nReceived a message: %s \n", d.Body)

			pFIle := &utils.PickFile{}

			err := json.Unmarshal(d.Body, pFIle)

			if err != nil {
				fmt.Printf("Error decoding JSON: %s", err)
			}
			if err := d.Ack(false); err != nil {
				fmt.Printf("Error acknowledging message : %s", err)
			} else {
				outpath := pFIle.Path[:len(pFIle.Path)-3] + "dxf"
				cmd := exec.Command("dwgread", pFIle.Path, "-O", "DXF", "-o", outpath)
				out, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("cmd.Run() failed with %s\n", err)
				}
				fmt.Println(string(out))
			}

		}
	}()

	// Stop for program termination
	<-stopChan

}
