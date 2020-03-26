package utils

import (
	"controller/config"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/streadway/amqp"
	tables "github.com/yossefaz/go_struct"
	"github.com/yossefaz/go_utils/queue"
	globalUtils "github.com/yossefaz/go_utils/utils"
)

func HandleError(err error, msg string, exit bool) {
	if err != nil {
		config.Logger.Log.Error(fmt.Sprintf("%s: %s", msg, err))
	}
	if exit {
		os.Exit(1)
	}
}

func MessageReceiver(m amqp.Delivery, rmq *queue.Rabbitmq) {
	switch m.Headers["From"] {
	case "Transformer":
		getMessageFromTransformer(m, rmq)
	case "Worker":
		GetMessageFromWorker(m, rmq)
	case "DAL":
		PoolReceiver(m, rmq)
	default:
		config.Logger.Log.Error("Received a message from a not known channel :", m.Headers["From"])
	}
}

func unpackFileMessage(m amqp.Delivery) *globalUtils.PickFile {
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
	log := config.Logger.Log
	if pFIle.Result["Transform"] == 1 {
		pFIle.Result = map[string]int{
			"BorderExist": 0,
			"InsideJer":   0,
		}

		mess, err := json.Marshal(pFIle)
		HandleError(err, "cannot convert transformed pFile to Json", false)
		res, err1 := rmq.SendMessage(mess, Constant.Channels.CheckDWG, Constant.Headers["CheckDWG"])
		HandleError(err1, "message sending error", false)
		config.Logger.Log.Info(res)
	} else if pFIle.Result["Transform"] == 0 {
		pFIle.Result["Transform"] = 1 // make the transform as error for create error
		createErrors := CreateDBMessage(map[string]interface{}{"check_status_id": pFIle.Id}, Constant.CRUD.CREATE, Constant.Cad_errors_table, fillStructFromResult(pFIle))
		mess := CreateDBMessage(map[string]interface{}{"Id": pFIle.Id}, Constant.CRUD.UPDATE, Constant.Cad_check_table, map[string]interface{}{"status_code": 20})
		sendMessageToQueue(createErrors, Constant.Channels.Dal_Req, Constant.Headers["Dal_Req"], rmq)
		sendMessageToQueue(mess, Constant.Channels.Dal_Req, Constant.Headers["Dal_Req"], rmq)
		log.Error("The transformer did not sucess to transform this file : ", pFIle.Path)
	}
}

func CheckResultsFromWorker(pFile *globalUtils.PickFile) int {
	for _, val := range pFile.Result {
		if globalUtils.Itob(val) {
			return 20
		}
	}
	return 10
}

func fillStructFromResult(pFile *globalUtils.PickFile) map[string]interface{} {
	resultMap := map[string]interface{}{}
	err := mapstructure.Decode(pFile.Result, &resultMap)
	HandleError(err, "cannot decode object to ORMKeyval", false)
	return resultMap
}

func CreateDBMessage(ids map[string]interface{}, crud string, table string, keyval map[string]interface{}) []byte {
	mess, err1 := json.Marshal(globalUtils.DbQuery{
		DbType:    Constant.DBType,
		Schema:    Constant.Schema,
		Table:     table,
		CrudT:     crud,
		Id:        ids,
		ORMKeyVal: keyval,
	})
	if err1 != nil {
		HandleError(err1, "cannot create a DB object from worker message", false)
	}
	return mess
}

func sendMessageToQueue(body []byte, queueName string, headers map[string]interface{}, rmq *queue.Rabbitmq) {
	message, err := rmq.SendMessage(body, queueName, headers)
	if err != nil {
		config.Logger.Log.Error(err)
	} else {
		config.Logger.Log.Info("SEND : " + message)
	}
}

func GetMessageFromWorker(m amqp.Delivery, rmq *queue.Rabbitmq) {
	pFIle := unpackFileMessage(m)
	status := CheckResultsFromWorker(pFIle)
	if pFIle.Status != status || status == 20 {
		createErrors := CreateDBMessage(map[string]interface{}{"check_status_id": pFIle.Id}, Constant.CRUD.CREATE, Constant.Cad_errors_table, fillStructFromResult(pFIle))
		mess := CreateDBMessage(map[string]interface{}{"Id": pFIle.Id}, Constant.CRUD.UPDATE, Constant.Cad_check_table, map[string]interface{}{"status_code": status})
		sendMessageToQueue(createErrors, Constant.Channels.Dal_Req, Constant.Headers["Dal_Req"], rmq)
		sendMessageToQueue(mess, Constant.Channels.Dal_Req, Constant.Headers["Dal_Req"], rmq)
	}
}

func PoolReceiver(m amqp.Delivery, rmq *queue.Rabbitmq) {
	switch m.Headers["Type"] {
	case "retrieve":
		getRetrieveResponse(m, rmq)
	case "update":
		getUpdateResponse(m)
	}
}

func getUpdateResponse(m amqp.Delivery) {
	config.Logger.Log.Info(string(m.Body))
}

func getRetrieveResponse(m amqp.Delivery, rmq *queue.Rabbitmq) {
	var res []tables.Cad_check_status
	err := json.Unmarshal(m.Body, &res)
	if err != nil {
		fmt.Println(err)
		config.Logger.Log.Error(err)
		HandleError(err, "MUST DISPATCH from POOL RECEIVER", false)
		config.Logger.Log.Error(string(m.Body))
	}
	for _, file := range res {
		message, err := json.Marshal(globalUtils.PickFile{
			Id:     file.ID,
			Path:   file.Path,
			Status: *file.Status_code,
			Result: map[string]int{
				"Transform": 0,
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
	mess := CreateDBMessage(map[string]interface{}{}, Constant.CRUD.RETRIEVE, Constant.Cad_check_table, map[string]interface{}{"status_code": 0})
	sendMessageToQueue(mess, Constant.Channels.Dal_Req, Constant.Headers["Dal_Req"], rmqConn)
}

func Scheduler(tick *time.Ticker, done chan bool, rmqConn *queue.Rabbitmq) {
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

func task(rmqConn *queue.Rabbitmq, t time.Time) {
	Pooling(rmqConn)
}
