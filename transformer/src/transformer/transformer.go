package main

import (
	"src/transformer/queue"
)

func main() {
	rmqConn := queue.NewRabbit("amqp://guest:guest@rabbitmq/", "transformDWG")
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.ListenMessage()
}
