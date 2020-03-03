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
func HandleError(err error, msg string) {
	log := config.Logger.Log
	if err != nil {
		log.Error("%s: %s", msg, err)
	}

}

func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	log := config.Logger.Log
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	HandleError(err, "Unable to convert message to json")
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
		HandleError(err, "cannot convert transformed pFile to Json")

		res, err1 := rmq.SendMessage(mess, "CheckDWG")
		HandleError(err1, "message sending error")
		config.Logger.Log.Info(res)

	} else if pFIle.Result["Transform"] == 0 {
		log.Error("The transformer did not sucess to transform this file : " , pFIle.Path)
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
		HandleError(err, "Cannot encode JSON")
		time.Sleep(time.Microsecond)

		res, err1 := rmqConn.SendMessage(message, "ConvertDWG")
		HandleError(err1, "message sending error")
		config.Logger.Log.Info(res)
	}
}