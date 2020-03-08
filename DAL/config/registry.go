package config

import (
	"dal/model"
)


var SchemaReg = map[string]model.Schema{}

func initReg() {
	SchemaReg["dwg_transformer"] = LocalConfig.DB.Mysql.Schema["dwg_transformer"]
}
func GetDBConf(schema string) model.Schema {
	return SchemaReg[schema]
}


