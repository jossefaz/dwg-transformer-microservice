package main

import (
	"github.com/tkanos/gonfig"
<<<<<<< HEAD:controller/config/config.go
	"github.com/yossefazoulay/go_utils/utils"
=======
>>>>>>> 87629761a850af82eefc4851d27a76e5f8d416cd:lib/config/config.go
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

func GetConfig(env string )  Configuration {
	configEnv["dev"] = "./config/config.dev.json"
	configEnv["prod"] = "./config/config.prod.json"
	configuration := Configuration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	HandleError(err, "Cannot load/read config file")
	return configuration
}

