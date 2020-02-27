package main

import (
	"listener/config"
	"listener/utils"
	"os"
	"github.com/yossefazoulay/go_utils/queue"

)

func main() {
	config.GetConfig(os.Args[1])
	rmqConn := queue.NewRabbit(config.LocalConfig.Queue.Rabbitmq.ConnString, config.LocalConfig.Queue.Rabbitmq.QueueNames)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(config.LocalConfig.Queue.Rabbitmq.Listennig, utils.MessageReceiver)
}
