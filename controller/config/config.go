package config

import (
	"github.com/tkanos/gonfig"
	"github.com/yossefazoulay/go_utils/utils"

)


type Configuration struct {
	Queue struct {
		Rabbitmq struct {
			ConnString string   `json:"ConnString"`
			QueueNames []string `json:"QueueNames"`
		} `json:"Rabbitmq"`
	} `json:"Queue"`
}

var configEnv = map[string]string{
	"dev" : "./config/config.dev.json",
	"prod" : "./config/config.prod.json",
}

var LocalConfig Configuration

func GetConfig(env string) {
	configuration := Configuration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	utils.HandleError(err, "Cannot load/read config file")
	LocalConfig = configuration
}

