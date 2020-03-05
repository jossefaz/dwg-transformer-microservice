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
		res:= execute(pFIle)
		mess, err1 :=rmq.SendMessage(res, config.LocalConfig.Queue.Rabbitmq.Result.Success)
		HandleError(err1, "message sending error", false)
		config.Logger.Log.Info(fmt.Sprintf(mess, err))
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
func execute(pfile *globalUtils.PickFile) []byte{
	resultConfig := getResultConfig()
	cmd := exec.Command("python", "bootstrap.py", pfile.Path, convertMapToString(pfile.Result))
	err := cmd.Run()
	if err != nil {
		HandleError(err, "cannot execute python", false)
		return setResult(pfile, pfile.Path, resultConfig.From, resultConfig.Fail, []int{0, 0})
	}
	return setResult(pfile, pfile.Path, resultConfig.From, resultConfig.Success, []int{1,1})
}

func setResult(pfile *globalUtils.PickFile, path string, from string, to string, result[]int)[]byte {

	keys := make([]string, 0, len(pfile.Result))
	for k := range pfile.Result {
		keys = append(keys, k)
	}
	mess, err := globalUtils.SetResultMessage(pfile, keys, result,  from, to, path)
	if err != nil {
		HandleError(err, "Cannot set output and cannot run command :" + err.Error() + err.Error(), false)
	}
	return mess
}
func getResultConfig() globalUtils.Result {
	return config.LocalConfig.Queue.Rabbitmq.Result
}

