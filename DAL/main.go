package main

import (
	"dal/config"
	"dal/log"
	"dal/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
)


func init(){
	environment, err := globalUtils.GetEnv("DEV_PROD")
	utils.HandleError(err, "Error while getting env variable", err != nil)
	config.GetConfig(environment)
}

func main() {

	log.GetLogger(os.Args[1])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)


}

