package main

import (
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
	"transformer/config"
	"transformer/utils"
)

func main() {

	config.GetConfig(os.Args[1], os.Args[2])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	globalUtils.HandleError(err, "Error Occured when RabbitMQ Init", &config.Logger)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)


}


