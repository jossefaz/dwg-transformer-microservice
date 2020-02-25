package utils

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"github.com/yossefazoulay/go_utils/queue"
	"os/exec"
)

func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	globalUtils.HandleError(err, "Error decoding message")
	if err := m.Ack(false); err != nil {
		fmt.Printf("Error acknowledging message : %s", err)
	} else {
		outpath := pFIle.Path[:len(pFIle.Path)-3] + "dxf"
		cmd := exec.Command("dwgread", pFIle.Path, "-O", "DXF", "-o", outpath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			rmq.SendMessage([]byte("DWG CONVERSION FAILED"), "ConvertDWG")
		}
		rmq.SendMessage([]byte("FILE CONVERTED SUCCESSFULLY"), "CheckDWG")
		fmt.Println(string(out))
	}
}
