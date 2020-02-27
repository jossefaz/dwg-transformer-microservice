package main

import (
	"controller/config"
	"controller/utils"
	"github.com/yossefazoulay/go_utils/queue"
	"os"
)

func main() {
	config.GetConfig(os.Args[1])
	rmqConn := queue.NewRabbit(config.LocalConfig.Queue.Rabbitmq.ConnString, config.LocalConfig.Queue.Rabbitmq.QueueNames)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(config.LocalConfig.Queue.Rabbitmq.Listennig, utils.MessageReceiver)
	utils.MockData(rmqConn)



}
