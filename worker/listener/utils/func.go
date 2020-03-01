package utils

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os/exec"
)

func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	if err := m.Ack(false); err != nil {
		fmt.Printf("Error acknowledging message : %s", err)
	} else {
		pFIle := &globalUtils.PickFile{}
		globalUtils.HandleError(
			json.Unmarshal(m.Body, pFIle), "Error decoding message in worker")
		cmd := exec.Command("python", "main.py", pFIle.Path)
		err = cmd.Run()
		if err != nil {
			rmq.SendMessage([]byte("WORKER WORKED"), "CheckedDWG")
		}
		rmq.SendMessage([]byte("WORKER WORKED"), "CheckedDWG")
	}
}
