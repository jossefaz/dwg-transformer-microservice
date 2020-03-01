package utils

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"time"
)



func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	globalUtils.HandleError(err, "test")
	if err := m.Ack(false); err != nil {
		fmt.Printf("Error acknowledging message : %s", err)
	}
	if pFIle.From == "Transformer" {
		getMessageFromTransformer(pFIle, rmq)
	} else if pFIle.From == "worker" {
		getMessageFromWorker(pFIle)
	}
}

func getMessageFromTransformer(pFIle *globalUtils.PickFile, rmq queue.Rabbitmq) {
	if pFIle.Result["Transform"] == 1 {
		pFIle.From = "controller"
		pFIle.Result = map[string]int{
			"BorderExist" : 0,
			"InsideJer" : 0,
		}
		mess, err := json.Marshal(pFIle)
		globalUtils.HandleError(err, "cannot convert transformed pFile to Json")
		rmq.SendMessage(mess, "CheckDWG")
	} else if pFIle.Result["Transform"] == 0 {
		fmt.Println("File :" + pFIle.Path + " failed to convert")
	}
}

func getMessageFromWorker(pFIle *globalUtils.PickFile) {
	fmt.Println("FROM WORKER :" + pFIle.Path + "")
}

func MockData(rmqConn queue.Rabbitmq) {
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
		globalUtils.HandleError(err, "Cannot encode JSON")
		time.Sleep(time.Microsecond)
		rmqConn.SendMessage(message, "ConvertDWG")
	}
}