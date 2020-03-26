package config

import (
	"dal/model"
	"fmt"

	"github.com/tkanos/gonfig"
	"github.com/yossefaz/go_utils/utils"
)

var LocalConfig Configuration

type Configuration struct {
	Queue struct {
		Rabbitmq struct {
			ConnString string   `json:"ConnString"`
			QueueNames []string `json:"QueueNames"`
			Listennig  []string `json:"Listennig"`
			Result     utils.Result
		} `json:"Rabbitmq"`
	} `json:"Queue"`
	DB struct {
		Mysql struct {
			Schema map[string]model.Schema
		} `json:"Mysql"`
	} `json:"DB"`
}

var configEnv = map[string]string{
	"dev":  "./config/config.dev.json",
	"prod": "./config/config.prod.json",
}

func GetConfig(env string) {
	configuration := Configuration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	if err != nil {
		fmt.Println("Cannot read config file")
	}
	LocalConfig = configuration
	initReg()
}
