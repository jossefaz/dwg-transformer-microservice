package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yossefazoulay/go_utils/queue"
	"os"
	"dal/config"
	"dal/utils"
)





func main() {

	config.GetConfig(os.Args[1])
	queueConf := config.LocalConfig.Queue.Rabbitmq
	rmqConn, err := queue.NewRabbit(queueConf.ConnString, queueConf.QueueNames)
	utils.HandleError(err, "Error Occured when RabbitMQ Init", err != nil)
	defer rmqConn.Conn.Close()
	defer rmqConn.ChanL.Close()
	rmqConn.OpenListening(queueConf.Listennig, utils.MessageReceiver)






	//fmt.Println(att)
	//fmt.Println("------------------------------")
	// get all records
	atts := []Attachements{} // a slice

	db.Where("status = ?", "0").Find(&atts)
	for _, v := range atts {
		fmt.Println("reference : ", v.Reference)
		fmt.Println("path : ", v.Path)
	}
	db.GetErrors()

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