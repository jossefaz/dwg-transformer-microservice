package utils

import (
	"controller/config"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"time"
)



func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	log := config.Logger.Log
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	globalUtils.HandleError(err, "Unable to convert message to json", &config.Logger)
	if err := m.Ack(false); err != nil {
		log.Error("Error acknowledging message : %s", err)
	}
	if pFIle.From == "Transformer" {
		getMessageFromTransformer(pFIle, rmq)
	} else if pFIle.From == "worker" {
		getMessageFromWorker(pFIle)
	}
}

func getMessageFromTransformer(pFIle *globalUtils.PickFile, rmq queue.Rabbitmq) {
	log:= config.Logger.Log
	if pFIle.Result["Transform"] == 1 {
		pFIle.From = "controller"
		pFIle.Result = map[string]int{
			"BorderExist" : 0,
			"InsideJer" : 0,
		}
		mess, err := json.Marshal(pFIle)
		globalUtils.HandleError(err, "cannot convert transformed pFile to Json", &config.Logger)
		rmq.SendMessage(mess, "CheckDWG")
	} else if pFIle.Result["Transform"] == 0 {
		log.Error("The transformer did not sucess to transfor this file : " , pFIle.Path)
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
		globalUtils.HandleError(err, "Cannot encode JSON", &config.Logger)
		time.Sleep(time.Microsecond)
		rmqConn.SendMessage(message, "ConvertDWG")
	}
}