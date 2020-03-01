package config

import (
	"github.com/tkanos/gonfig"
	"github.com/yossefazoulay/go_utils/utils"
	"github.com/yossefazoulay/go_utils/logs"
)

var LocalConfig Configuration
var Logger logs.Logger


type Configuration struct {
	Queue struct {
		Rabbitmq struct {
			ConnString string   `json:"ConnString"`
			QueueNames []string `json:"QueueNames"`
			Listennig  []string `json:"Listennig"`
			Result    utils.Result
		} `json:"Rabbitmq"`
	} `json:"Queue"`
	OutputFormat   string
	FileExtensions map[string]string
	Logs struct {
		Main struct {
			Path  string
			Level string
		}
	}
}

var configEnv = map[string]string{
	"dev" : "./config/config.dev.json",
	"prod" : "./config/config.prod.json",
}


func GetConfig(env string, output string ) {
	configuration := Configuration{}
	err := gonfig.GetConf(configEnv[env], &configuration)
	utils.HandleError(err, "Cannot load/read config file")
	LocalConfig = configuration
	LocalConfig.OutputFormat = output
	Logger = logs.InitLogs(LocalConfig.Logs.Main.Path, LocalConfig.Logs.Main.Level)
}


