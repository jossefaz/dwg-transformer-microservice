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

var configEnv = make(map[string]string)

func GetConfig(env string )  Configuration {
	configEnv["dev"] = "./config/config.dev.json"
	configEnv["prod"] = "./config/config.prod.json"
	configuration := Configuration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	utils.HandleError(err, "Cannot load/read config file")
	return configuration
}

