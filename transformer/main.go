package main

import (
	"os"
	"transformer/config"
	"transformer/utils"
	"github.com/yossefazoulay/go_utils/queue"

)

func main() {

	config.GetConfig(os.Args[1], os.Args[2])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)

}


