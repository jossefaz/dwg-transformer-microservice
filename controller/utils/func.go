package utils

import (
	"controller/config"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
	"time"
)
func HandleError(err error, msg string, exit bool) {
	if err != nil {
		config.Logger.Log.Error(fmt.Sprintf("%s: %s", msg, err))
	}
	if exit {
		os.Exit(1)
	}
}

func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	log := config.Logger.Log
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	HandleError(err, "Unable to convert message to json", false)
	if err := m.Ack(false); err != nil {
		log.Error("Error acknowledging message : %s", err)
	}
	if pFIle.From == "Transformer" {
		getMessageFromTransformer(pFIle, rmq)
	} else if pFIle.From == "Worker" {
		getMessageFromWorker(pFIle)
	}
}

func getMessageFromTransformer(pFIle *globalUtils.PickFile, rmq queue.Rabbitmq) {
	log:= config.Logger.Log
	if pFIle.Result["Transform"] == 1 {
		pFIle.From = "Controller"
		pFIle.Result = map[string]int{
			"BorderExist" : 0,
			"InsideJer" : 0,
		}
		pFIle.To = "CheckDWG"



		mess, err := json.Marshal(pFIle)
		HandleError(err, "cannot convert transformed pFile to Json", false)

		res, err1 := rmq.SendMessage(mess, "CheckDWG")
		HandleError(err1, "message sending error", false)

		config.Logger.Log.Info(res)

	} else if pFIle.Result["Transform"] == 0 {
		log.Error("The transformer did not sucess to transform this file : " , pFIle.Path)
	}
}

func getMessageFromWorker(pFIle *globalUtils.PickFile) {
	config.Logger.Log.Info("FROM WORKER :" + pFIle.Path + "")
}

func PoolReceiver(m amqp.Delivery, rmq queue.Rabbitmq) {
	type Timestamp time.Time
	type attachements struct {
		Reference int
		Status int
		StatusDate Timestamp
		Path string
	}
	log := config.Logger.Log
	res := []attachements{}
	mess := json.Unmarshal(m.Body, res)
	if err := m.Ack(false); err != nil {
		log.Error("Error acknowledging message : %s", err)
	}
	for _, file := range res {
		message, err := json.Marshal(globalUtils.PickFile{
			Path: file.Path,
			Result : map[string]int{
				"Transform" : 0,
			},
			From : "controller",
			To : "ConvertDWG",
		})
		HandleError(err, "Cannot encode JSON", false)
		time.Sleep(time.Microsecond)
		res, err1 := rmq.SendMessage(message, "ConvertDWG")
		HandleError(err1, "message sending error", false)
		config.Logger.Log.Info(res)
	}
	fmt.Println(mess)
}

func Pooling(rmqConn queue.Rabbitmq) {
	mess, _ := json.Marshal(globalUtils.DbQuery{
		Schema:"dwg_transformer",
		Table:  "Attachments",
		CrudT:  "retrieve",
		Id: []int{},
		ORMKeyVal: map[string]interface{}{
			"status" : 0,
		},
	})
	rmqConn.SendMessage(mess, "Dal_Req")
	}
