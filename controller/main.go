package main

import (
	"controller/config"
	"controller/utils"
	"github.com/yossefazoulay/go_utils/queue"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.GetConfig(os.Args[1])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()

	tick := time.NewTicker(time.Second * 10)

	done := make(chan bool)
	go scheduler(tick, done, rmqConn)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	rmqConn.ListenMessage(utils.PoolReceiver, queueConf.Listennig[2])
	rmqConn.ListenMessage(utils.MessageReceiver, queueConf.Listennig[0])
	rmqConn.ListenMessage(utils.MessageReceiver, queueConf.Listennig[1])

	<-sigs
	done <- true

}

func scheduler(tick *time.Ticker, done chan bool, rmqConn queue.Rabbitmq) {
	task(rmqConn, time.Now())
	for {
		select {
		case t := <-tick.C:
			task(rmqConn, t)
		case <-done:
			return
		}
	}
}

func task(rmqConn queue.Rabbitmq, t time.Time) {
	utils.Pooling(rmqConn)
}
