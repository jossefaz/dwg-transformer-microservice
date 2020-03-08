package utils

import (
	"dal/config"
	"dal/model"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
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
	dbQ := unpackMessage(m)
	dbconf := config.GetDBConf(dbQ.Schema)
	db := model.ConnectToDb(dbconf.Dialect, dbconf.ConnString)
	res := dispatcher(db, dbQ)
	rmq.SendMessage(res, "Dal_Res")
	defer db.Close()
}

func dispatcher(db *model.CDb, dbQ *globalUtils.DbQuery ) []byte {
	switch dbQ.CrudT {
	case "retrieve":
		res := db.Retrieve(dbQ)
		return res
	case "update":
		db.Update(dbQ)
		return []byte{}
	default:
		config.Logger.Log.Error("CRUD operation must be one of the following : retrieve, update | delete and create not supported yet")
		return []byte{}
	}

}


func unpackMessage(m amqp.Delivery) *globalUtils.DbQuery {
	dbQ := &globalUtils.DbQuery{}
	err := json.Unmarshal(m.Body, dbQ)
	if err := m.Ack(false); err != nil {
		config.Logger.Log.Error("Error acknowledging message : %s", err)
	}
	HandleError(err, "Error decoding DB message", false)
	return dbQ
}


//"mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local"