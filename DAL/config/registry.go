package config

import (
	"dal/model"
)

var TablesReg = map[string]interface{}{
	"Attachments" : []model.Attachements{},
}

var SchemaReg = map[string]model.Schema{}

func initReg() {
	SchemaReg["dwg_transformer"] = LocalConfig.DB.Mysql.Schema["dwg_transformer"]
}
func GetDBConf(schema string) model.Schema {
	return SchemaReg[schema]
}
func GetTableStruct(tablename string) interface{} {
	return TablesReg[tablename]
}

