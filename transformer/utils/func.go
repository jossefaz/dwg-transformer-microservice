package utils

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/yossefazoulay/go_utils/queue"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os/exec"
	"strings"
	"transformer/config"
)

func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	resultConfig := getResultConfig()
	log := config.Logger.Log
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	globalUtils.HandleError(err, "Error decoding message", config.Logger)
	if pFIle.From !=  resultConfig.From{
		if err := m.Ack(false); err != nil {
			log.Error("Error acknowledging message : %s", err)
		} else {
			res:= execute(pFIle, config.LocalConfig.OutputFormat)
			rmq.SendMessage(res, resultConfig.Success)
		}
	}
}

func getResultConfig() globalUtils.Result {
	return config.LocalConfig.Queue.Rabbitmq.Result
}

func execute(pfile *globalUtils.PickFile, output string) []byte{
	resultConfig := getResultConfig()
	outpath := getOutputPath(pfile.Path, output)
	cmd := exec.Command("dwgread", pfile.Path, "-O", output, "-o", outpath)
	err := cmd.Run()
	if err != nil {
		return setResult(pfile, pfile.Path, resultConfig.From, true)
	}
	return setResult(pfile, outpath, resultConfig.From, false)
}

func setResult(pfile *globalUtils.PickFile, path string, from string, error bool)[]byte {
	execRes := 1
	if error {
		execRes = 0
	}
	mess, err := globalUtils.SetResultMessage(pfile, []string{"Transform"}, []int {execRes},  from, path)
	if err != nil {
		globalUtils.HandleError(err, "Cannot set output and cannot run command :" + err.Error() + err.Error(), config.Logger)
	}
	return mess
	}


func getOutputPath(basePath string, output string) string {
	fileExt := config.LocalConfig.FileExtensions[output]
	outpath := strings.Split(basePath, ".")[0] + fileExt
	return outpath
}
