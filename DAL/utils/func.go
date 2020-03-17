package utils

import (
	"dal/config"
	"dal/log"
	"dal/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
)

func HandleError(err error, msg string, exit bool) {
	if err != nil {
		log.Logger.Log.Error(fmt.Sprintf("%s: %s", msg, err))
	}
	if exit {
		os.Exit(1)
	}
}




func MessageReceiver(m amqp.Delivery, rmq *queue.Rabbitmq)  {
	dbQ, err := unpackMessage(m)
	HandleError(err, "cannot unpack message received in DAL to DB", err!=nil)
	dbconf := config.GetDBConf(dbQ.DbType, dbQ.Schema)
	db, err := model.ConnectToDb(dbconf.Dialect, dbconf.ConnString)
	HandleError(err, "cannot connect to DB", err!=nil)
	res, err := dispatcher(db, dbQ)
	if err != nil {
		log.Logger.Log.Error(err)
	} else {
		rmq.SendMessage(res, "Dal_Res", map[string]interface{}{
			"From" : "DAL",
			"To" : "Dal_Res",
			"Type" : dbQ.CrudT,
		})
	}
	defer db.Close()
}

func dispatcher(db *model.CDb, dbQ *globalUtils.DbQuery ) ([]byte, error) {
	switch dbQ.CrudT {
	case "retrieve":
		res, err := db.RetrieveRow(dbQ)
		if err != nil {
			return nil, err
		}
		return res, nil
	case "update":
		res, err := db.UpdateRow(dbQ)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, errors.New("CRUD operation must be one of the following : retrieve, update | delete and create not supported yet")
	}

}


func unpackMessage(m amqp.Delivery) (*globalUtils.DbQuery, error) {
	dbQ := &globalUtils.DbQuery{}
	err := json.Unmarshal(m.Body, dbQ)
	if err := m.Ack(false); err != nil {
		log.Logger.Log.Error("Error acknowledging message : %s", err)
		return &globalUtils.DbQuery{}, err
	}
	HandleError(err, "Error decoding DB message", false)
	return dbQ, nil
}


//"mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local"