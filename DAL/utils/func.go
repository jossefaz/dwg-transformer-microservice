package utils

import (
	"dal/config"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
)

type cDb struct {
	*gorm.DB
}


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
	db := ConnectToDb(dbconf.Dialect, dbconf.ConnString)

	defer db.Close()
}

func unpackMessage(m amqp.Delivery) *globalUtils.DbQuery {
	dbQ := &globalUtils.DbQuery{}
	err := json.Unmarshal(m.Body, dbQ)
	HandleError(err, "Error decoding DB message", false)
	return dbQ
}

func ConnectToDb(dialect string, connString string) *cDb {
	db, err := gorm.Open(dialect, connString)
	HandleError(err, "Error connecting to db", err != nil)
	db.DB()
	db.DB().Ping()
	var dup = cDb{ db}
	return &dup
}

func (db cDb) Retrieve(tableName string, query string, val string) {
	atts :=  config.GetTableStruct(tableName)
	db.Where(query, val).Find(&atts)
	db.GetErrors()
}
//"mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local"