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

func MessageReceiver(m amqp.Delivery, rmq *queue.Rabbitmq)  {
	if err := m.Ack(false); err != nil {
		config.Logger.Log.Error(fmt.Sprintf("Error acknowledging message : %s", err))
	} else {
		pFIle := &globalUtils.PickFile{}
		HandleError(json.Unmarshal(m.Body, pFIle), "Error decoding message in worker",false)
		res:= execute(pFIle)
		mess, err1 :=rmq.SendMessage(res, config.LocalConfig.Queue.Rabbitmq.Result.Success, config.LocalConfig.Queue.Rabbitmq.Result.From)
		HandleError(err1, "message sending error", false)
		config.Logger.Log.Info(fmt.Sprintf(mess, err))
	}
}



func convertMapKeysToString(customMap map[string]int) string {
	var b bytes.Buffer

	for k := range customMap {
		s := fmt.Sprintf("%s ", k)
		b.WriteString(s)
	}
	return b.String()

}
func execute(pfile *globalUtils.PickFile) []byte{
	cmd := exec.Command("python", "bootstrap.py", pfile.Path, convertMapKeysToString(pfile.Result))
	out, err := cmd.CombinedOutput()
	if err != nil {
		HandleError(err, "cannot execute python", false)
	}
	return setResult(pfile, out)
}

func setResult(pfile *globalUtils.PickFile, cmdRes []byte)[]byte {
	resMap := make(map[string]int)
	err := json.Unmarshal(cmdRes, &resMap)
	for k, v := range resMap {
		pfile.Result[k] = v
	}
	res, err := json.Marshal(pfile)
	if err != nil {
		HandleError(err, "Cannot set output and cannot run command :" + err.Error() + err.Error(), false)
	}
	return res
}
func getResultConfig() globalUtils.Result {
	return config.LocalConfig.Queue.Rabbitmq.Result
}

