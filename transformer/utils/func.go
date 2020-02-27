package utils

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"github.com/yossefazoulay/go_utils/queue"
	"os/exec"
	"strings"
	"transformer/config"
)

func MessageReceiver(m amqp.Delivery, rmq queue.Rabbitmq)  {
	pFIle := &globalUtils.PickFile{}
	err := json.Unmarshal(m.Body, pFIle)
	globalUtils.HandleError(err, "Error decoding message")
	if err := m.Ack(false); err != nil {
		fmt.Printf("Error acknowledging message : %s", err)
	} else {
		resultConfig := getResultConfig()
		res:= execute(pFIle, config.LocalConfig.OutputFormat)
		rmq.SendMessage(res, resultConfig.Success)
	}
}

func getResultConfig() globalUtils.Result {
	return config.LocalConfig.Queue.Rabbitmq.Result
}

func execute(pfile *globalUtils.PickFile, output string) []byte{
	resultConfig := getResultConfig()
	outpath := getOutputPath(pfile.Path, output)
	cmd := exec.Command("dwgread", pfile.Path, "-O", output, "-o", outpath)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return globalUtils.SetResultMessage(pfile, []string{"Transform"}, []int {0},  resultConfig.From, pfile.Path)
	}
	return globalUtils.SetResultMessage(pfile, []string{"Transform"}, []int {1},  resultConfig.From, outpath)
}

func getOutputPath(basePath string, output string) string {
	fileExt := config.LocalConfig.FileExtensions[output]
	outpath := strings.Split(basePath, ".")[0] + fileExt
	return outpath
}
