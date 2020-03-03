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
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init")
	if err == nil {
		defer rmqConn.Conn.Close()
		defer rmqConn.ChanL.Close()
		utils.MockData(rmqConn)
		rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)
	}

}
