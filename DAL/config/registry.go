package config

import (
	"dal/model"
)

var SchemaReg = map[string]map[string]model.Schema{}

func initReg() {
	SchemaReg["mysql"] = map[string]model.Schema{
		"dwg_transformer": LocalConfig.DB.Mysql.Schema["dwg_transformer"],
	}

}
func GetDBConf(dbtype string, schema string) model.Schema {
	return SchemaReg[dbtype][schema]
}
