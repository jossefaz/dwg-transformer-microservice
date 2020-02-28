package main

import (
	"controller/config"
	"controller/utils"
	"encoding/json"
	"fmt"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
	"time"
)

func main() {
	config.GetConfig(os.Args[1])
	rmqConn := queue.NewRabbit(config.LocalConfig.Queue.Rabbitmq.ConnString, config.LocalConfig.Queue.Rabbitmq.QueueNames)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	fmt.Println(rmqConn.Conn.ConnectionState())
	root := "./shared"
	files := globalUtils.ListFilesInDir(root)
	for _, file := range files {
		message, err := json.Marshal(globalUtils.PickFile{
			Path: file,
			Result : map[string]int{
				"Transform" : 0,
			},
			From : "controller",
		})
		globalUtils.HandleError(err, "Cannot decode JSON")
		time.Sleep(time.Second)
		rmqConn.SendMessage(message, "ConvertDWG")
	}
	rmqConn.ListenMessage(utils.MessageReceiver, config.LocalConfig.Queue.Rabbitmq.Listennig[0])
	rmqConn.ListenMessage(utils.MessageReceiver, config.LocalConfig.Queue.Rabbitmq.Listennig[1])
	utils.MockData(rmqConn)



}
