package main

import (
	"os"
	"dwg.transformer/main/lib/config"
	"dwg.transformer/main/lib/queue"
)

func main() {
	configuration := config.GetConfig(os.Args[1])
	rmqConn := queue.NewRabbit(configuration.Queue.Rabbitmq.ConnString, configuration.Queue.Rabbitmq.QueueNames.ConvertDWG)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.ListenMessage()
}
