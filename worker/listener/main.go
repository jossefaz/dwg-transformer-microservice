package main

import (
	"github.com/yossefazoulay/go_utils/queue"
	"os"
	"listener/config"
	"listener/utils"
)

func main() {
	config.GetConfig(os.Args[1])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)
}
