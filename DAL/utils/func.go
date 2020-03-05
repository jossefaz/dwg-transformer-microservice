package utils

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
	"dal/config"
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

	dbQ := &globalUtils.DbQuery{}
	dbconf := config.GetDBConf(dbQ.Schema)
	err := json.Unmarshal(m.Body, dbQ)
	HandleError(err, "Error decoding message", false)
	db, errdb := connectToDb(dbconf.Dialect, dbconf.ConnString)
	HandleError(errdb, "Error decoding message", errdb != nil)
	defer db.Close()
	log.Info(dbQ.ORMSQL)
}

func connectToDb(dialect string, connString string) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, connString)
	if err!=nil {
		fmt.Println("Cannot connect to DB", err)
		return &gorm.DB{}, err
	}
	db.DB()
	db.DB().Ping()
	return db, nil
}
//"mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local"