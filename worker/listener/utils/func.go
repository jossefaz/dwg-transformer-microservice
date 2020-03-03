package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"listener/config"
	"os"
	"os/exec"
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
	if err := m.Ack(false); err != nil {
		config.Logger.Log.Error(fmt.Sprintf("Error acknowledging message : %s", err))
	} else {
		pFIle := &globalUtils.PickFile{}
		HandleError(json.Unmarshal(m.Body, pFIle), "Error decoding message in worker",false)
		cmd := exec.Command("python", "main.py", pFIle.Path, convertMapToString(pFIle.Result))
		err = cmd.Run()
		if err != nil {
			mess, err1 :=rmq.SendMessage([]byte("WORKER DOES NOT WORKED"), "CheckedDWG")
			HandleError(err1, "message sending error", false)
			config.Logger.Log.Error(fmt.Sprintf(mess, err))
		} else {
			mess, err1 :=rmq.SendMessage([]byte("WORKER WORKED"), "CheckedDWG")
			HandleError(err1, "message sending error", false)
			config.Logger.Log.Info(mess)
		}


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
