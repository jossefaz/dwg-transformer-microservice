package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"listener/config"
	"os/exec"
)

func HandleError(err error, msg string) {
	if err != nil {
		config.Logger.Log.Error("%s: %s", msg, err)
	}
}

func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	if err := m.Ack(false); err != nil {
		fmt.Printf("Error acknowledging message : %s", err)
	} else {
		pFIle := &globalUtils.PickFile{}
		globalUtils.HandleError(
			json.Unmarshal(m.Body, pFIle), "Error decoding message in worker", &config.Logger)
		cmd := exec.Command("python", "main.py", pFIle.Path, convertMapToString(pFIle.Result))
		err = cmd.Run()
		if err != nil {
			rmq.SendMessage([]byte("WORKER DOES NOT WORKED"), "CheckedDWG")
		}
		rmq.SendMessage([]byte("WORKER WORKED"), "CheckedDWG")
	}
}

func convertMapToString(customMap map[string]int) string {
	var b bytes.Buffer

	for k,v := range customMap {
		s := fmt.Sprintf("%s=\"%v\"", k, v)
		b.WriteString(s)
	}
	return b.String()

}

func convertMapKeysToString(customMap map[string]int) string {
	var b bytes.Buffer

	for k := range customMap {
		s := fmt.Sprintf("%s ", k)
		b.WriteString(s)
	}
	return b.String()

}
