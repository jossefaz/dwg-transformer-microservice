package main

import (
	"encoding/json"
	"github.com/yossefazoulay/go_utils/queue"
	"github.com/yossefazoulay/go_utils/utils"
	"controller/config"
	"os"
	"time"
)

func main() {
	configuration := config.GetConfig(os.Args[1])
	rmqConn := queue.NewRabbit(configuration.Queue.Rabbitmq.ConnString, configuration.Queue.Rabbitmq.QueueNames.ConvertDWG)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	root := "./"
	files := utils.ListFilesInDir(root)

	for _, file := range files {

		message, err := json.Marshal(utils.PickFile{
			Name: "File Uploaded",
			Path: file,
		})
		utils.HandleError(err, "Cannot decode JSON")
		time.Sleep(time.Second)

		rmqConn.SendMessage(message)
	}
	rmqConn.ListenMessage()
}
