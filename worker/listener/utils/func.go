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
		err := json.Unmarshal(m.Body, pFIle)
		cmd := exec.Command("python", "main.py", pFIle.Path)
		out, err := cmd.CombinedOutput()
		if err != nil {
			globalUtils.HandleError(err, "Error decoding message")
		}
		fmt.Println(string(out))
	}
}
