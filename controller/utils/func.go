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

func MessageReceiver(m amqp.Delivery, rmq *queue.Rabbitmq)  {
	switch m.Headers["From"] {
	case "Transformer":
		getMessageFromTransformer(m, rmq)
	case "Worker" :
		getMessageFromWorker(m, rmq)
	case "DAL":
		PoolReceiver(m, rmq)
	default:
		config.Logger.Log.Error("Received a message from a not known channel :", m.Headers["From"])
	}
}

func unpackFileMessage(m amqp.Delivery) *globalUtils.PickFile{
	log := config.Logger.Log
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	HandleError(err, "Unable to convert message to json", false)
	if err := m.Ack(false); err != nil {
		log.Error("Error acknowledging message : %s", err)
	}
	return pFIle
}

func getMessageFromTransformer(m amqp.Delivery, rmq *queue.Rabbitmq) {
	pFIle := unpackFileMessage(m)
	log:= config.Logger.Log
	if pFIle.Result["Transform"] == 1 {
		pFIle.Result = map[string]int{
			"BorderExist" : 0,
			"InsideJer" : 0,
		}

		mess, err := json.Marshal(pFIle)
		HandleError(err, "cannot convert transformed pFile to Json", false)
		res, err1 := rmq.SendMessage(mess, Constant.Channels.CheckDWG,  Constant.Headers["CheckDWG"])
		HandleError(err1, "message sending error", false)
		config.Logger.Log.Info(res)
	} else if pFIle.Result["Transform"] == 0 {
		log.Error("The transformer did not sucess to transform this file : " , pFIle.Path)
	}
}

func getMessageFromWorker(m amqp.Delivery, rmq *queue.Rabbitmq) {
	pFIle := unpackFileMessage(m)

	mess, err := json.Marshal(globalUtils.DbQuery{
		DbType: "mysql",
		Schema:"dwg_transformer",
		Table:  "Attachments",
		CrudT:  "update",
		Id: map[string]interface{}{
			"reference" : pFIle.Id,
		},
		ORMKeyVal: map[string]interface{}{
			"status" : 1,
		},
	})
	if err != nil {
		fmt.Println(err)
		config.Logger.Log.Error(err)
		HandleError(err, "cannot unmarshal json from worker", false)
		config.Logger.Log.Error(string(m.Body))
	}
	message, err := rmq.SendMessage(mess, Constant.Channels.Dal_Req, Constant.Headers["Dal_Req"])
	if err != nil {
		config.Logger.Log.Error(err)
	} else {
		config.Logger.Log.Info("SEND : " + message)
	}

}

func PoolReceiver(m amqp.Delivery, rmq *queue.Rabbitmq) {
	switch m.Headers["Type"] {
	case "retrieve":
		getRestrieveResponse(m, rmq)
	case "update":
		getUpdateResponse(m)
	}
}

func getUpdateResponse(m amqp.Delivery){
	config.Logger.Log.Info(string(m.Body))
}

func getRestrieveResponse(m amqp.Delivery, rmq *queue.Rabbitmq){
	type Timestamp time.Time
	type Attachements struct {
		Reference int
		Status int
		StatusDate Timestamp
		Path string
	}

	var res []Attachements
	err := json.Unmarshal(m.Body, &res)
	if err != nil {
		fmt.Println(err)
		config.Logger.Log.Error(err)
		HandleError(err, "MUST DISPATCH from POOL RECEIVER", false)
		config.Logger.Log.Error(string(m.Body))
	}
	for _, file := range res {
		message, err := json.Marshal(globalUtils.PickFile{
			Id: file.Reference,
			Path: file.Path,
			Result : map[string]int{
				"Transform" : 0,
			},
		})
		HandleError(err, "Cannot encode JSON", false)
		time.Sleep(time.Microsecond)
		res, err1 := rmq.SendMessage(message, Constant.Channels.ConvertDWG, Constant.Headers["ConvertDWG"])
		HandleError(err1, "message sending error", false)
		config.Logger.Log.Info(res)
	}

}

func Pooling(rmqConn *queue.Rabbitmq) {
	mess, _ := json.Marshal(globalUtils.DbQuery{
		DbType: "mysql",
		Schema:"dwg_transformer",
		Table:  "Attachments",
		CrudT:  "retrieve",
		Id: map[string]interface{}{},
		ORMKeyVal: map[string]interface{}{
			"status" : 0,
		},
	})
	rmqConn.SendMessage(mess, Constant.Channels.Dal_Req, Constant.Headers["Dal_Req"])
}
