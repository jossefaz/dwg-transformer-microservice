package main

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Queue struct {
		Rabbitmq struct {
			ConnString string `json:"ConnString"`
			QueueNames struct {
				ConvertDWG string `json:"ConvertDWG"`
			} `json:"QueueNames"`
		} `json:"Rabbitmq"`
	} `json:"Queue"`
}

var configEnv = make(map[string]string)

func GetConfig(env string)  Configuration {
	configEnv["dev"] = "./config/files/config.dev.json"
	configEnv["prod"] = "./config/files/config.prod.json"
	configuration := Configuration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	HandleError(err, "Cannot load/read config file")
	return configuration
}

