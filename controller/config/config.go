package config

import (
	"github.com/tkanos/gonfig"
	"github.com/yossefazoulay/go_utils/utils"
	"github.com/yossefazoulay/go_utils/logs"
)


type Configuration struct {
	Queue struct {
		Rabbitmq struct {
			ConnString string   `json:"ConnString"`
			QueueNames []string `json:"QueueNames"`
			Listennig  []string `json:"Listennig"`
		} `json:"Rabbitmq"`
	} `json:"Queue"`
	Logs struct{
	Main struct{
		Path string
		Level string
	}
}
}

var configEnv = map[string]string{
	"dev" : "./config/config.dev.json",
	"prod" : "./config/config.prod.json",
}

var LocalConfig Configuration
var Logger logs.Logger

func GetConfig(env string) {
	configuration := Configuration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	utils.HandleError(err, "Cannot load/read config file")
	LocalConfig = configuration
	Logger = logs.InitLogs(LocalConfig.Logs.Main.Path, LocalConfig.Logs.Main.Level)
}

