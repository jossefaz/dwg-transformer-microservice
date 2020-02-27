package utils

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"time"
)



func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	globalUtils.HandleError(err, "test")
	fmt.Println(pFIle)
	if pFIle.From == "Transformer" {
		if pFIle.Result["Transform"] == 1 {
			pFIle.From = "controller"
			pFIle.Result = map[string]int{
				"BorderExist" : 0,
				"InsideJer" : 0,
			}
			mess, err := json.Marshal(pFIle)
			globalUtils.HandleError(err, "cannot convert transformed pFile to Json")
			rmq.SendMessage(mess, "CheckDWG")
		} else if pFIle.Result["Transform"] == 0 {
			fmt.Println("File :" + pFIle.Path + " failed to convert")
		}
	} else if pFIle.From == "worker" {
		fmt.Println("FROM WORKER :" + pFIle.Path + "")
	}

	//globalUtils.HandleError(err, "Error decoding message")
	//if err := m.Ack(false); err != nil {
	//	fmt.Printf("Error acknowledging message : %s", err)
	//} else {
	//	resultConfig := getResultConfig()
	//	res:= execute(pFIle, config.LocalConfig.OutputFormat)
	//	rmq.SendMessage(res, resultConfig.Success)
	//}
}

func getMessageFromTransformer() {

}

func MockData(rmqConn queue.Rabbitmq) {
	root := "./"
	files := globalUtils.ListFilesInDir(root)
	for _, file := range files {
		message, err := json.Marshal(globalUtils.PickFile{
			Path: file,
			Result : map[string]int{
				"Transform" : 0,
			},
			From : "controller",
		})
		globalUtils.HandleError(err, "Cannot decode JSON")
		time.Sleep(time.Second)
		rmqConn.SendMessage(message, "ConvertDWG")
	}
}