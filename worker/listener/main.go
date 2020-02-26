package main

import (
	"listener/config"
	"listener/utils"
	"os"
	"github.com/yossefazoulay/go_utils/queue"

)

func main() {
	configuration := config.GetConfig(os.Args[1])
	rmqConn := queue.NewRabbit(configuration.Queue.Rabbitmq.ConnString, configuration.Queue.Rabbitmq.QueueNames)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.ListenMessage(utils.MessageReceiver, "CheckDWG")
}