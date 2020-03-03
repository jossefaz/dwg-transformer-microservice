package main

import (
	"controller/config"
	"os"
)

func main() {

	config.GetConfig(os.Args[1])
	config.Logger.Log.Info("Ceci est un test")
	//queueConf := config.LocalConfig.Queue.Rabbitmq
	//rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	//globalUtils.HandleError(err, "Error Occured when RabbitMQ Init", config.Logger)
	//defer rmqConn.Conn.Close()
	//defer rmqConn.ChanL.Close()
	//utils.MockData(rmqConn)
	//rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)
}
