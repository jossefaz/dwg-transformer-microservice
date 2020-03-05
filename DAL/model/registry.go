package model

import "dal/config"

var Tables = map[string]interface{}{
	"Attachments" : Attachements{},
}

var Schema = map[string] config.Schema {
	"dwg_transformer" : config.LocalConfig.DB.Mysql.Schema["dwg_transformer"],
}

