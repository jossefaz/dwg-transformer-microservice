package main

import (
	"github.com/yossefaz/dwg-transformer-microservice/controller/config"
	"controller/utils"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/yossefaz/go_utils/queue"
	globalUtils "github.com/yossefaz/go_utils/utils"
)

func init() {
	environment, err := globalUtils.GetEnv("DEV_PROD")
	utils.HandleError(err, "Error while getting env variable", err != nil)
	config.GetConfig(environment)
}

func main() {

	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)

	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()

	poolInterval, err := strconv.Atoi(os.Getenv("POOL_INTERVAL"))
	utils.HandleError(err, "Cannot set environment variable POOL_INTERVAL to integer, check if POOL_INTERVAL is set and if it is an integer", err != nil)
	tick := time.NewTicker(time.Second * time.Duration(poolInterval))
	done := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)
	go utils.Scheduler(tick, done, rmqConn)
	<-sigs
	done <- true

}
