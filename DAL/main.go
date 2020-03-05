package main

import (
	"dal/config"
	"dal/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)





func main() {

	config.GetConfig(os.Args[1])
	//queueConf := config.LocalConfig.Queue.Rabbitmq
	//rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	//utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	//defer rmqConn.Conn.Close()
	//defer rmqConn.ChanL.Close()
	//rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)
	dbconf := config.GetDBConf("dwg_transformer")
	db := utils.ConnectToDb(dbconf.Dialect, dbconf.ConnString)
	db.Retrieve("Attachments", "status = ?", "0")








	//fmt.Println(att)
	//fmt.Println("------------------------------")
	// get all records


	//dbConn.Find(&activities)
	//fmt.Println(activities)

	//rows, err := db.Model(&Attachements{}).Rows()
	//defer rows.Close()
	//if err != nil {
	//	panic(err)
	//}
	//for rows.Next() {
	//	db.ScanRows(rows, &att)
	//	fmt.Println(att)
	//}


}