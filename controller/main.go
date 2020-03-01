package main

import (
	"controller/config"
	"controller/utils"
	"github.com/yossefazoulay/go_utils/queue"
	"os"
)

func main() {
	config.GetConfig(os.Args[1])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)
	utils.MockData(rmqConn)



}
