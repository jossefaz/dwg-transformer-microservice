package utils

import (
	"os"
)

type channels struct {
	CheckDWG string
	CheckedDWG string
	ConvertDWG string
	ConvertedDWG string
	Dal_Req string
	Dal_Res string

}

type crud struct {
	CREATE string
	RETRIEVE string
	UPDATE string
	DELETE string
}

type headers map[string]map[string]interface{}

type constants struct {
	CheckDWG string
	From string
	Channels channels
	Headers headers
	Cad_check_table string
	Cad_errors_table string
	DBType string
	Schema string
	CRUD crud

}

var Constant = constants{
	CheckDWG: "CheckDWG",
	From:     "Controller",
	Channels: channels{
		     "CheckDWG",
		   "CheckedDWG",
		   "ConvertDWG",
		 "ConvertedDWG",
		      "Dal_Req",
		      "Dal_Res",
	},
	Headers: headers{
		"ConvertDWG" : map[string]interface{}{
			"From": "Controller",
			"To":   "ConvertDWG",
		},
		"Dal_Req" : map[string]interface{}{
			"From": "Controller",
			"To":   "Dal_Req",
		},
	},
	Cad_check_table : os.Getenv("CAD_STATUS"),
	Cad_errors_table: os.Getenv("CAD_ERRORS"),
	DBType: os.Getenv("DB"),
	Schema: os.Getenv("SCHEMA"),
	CRUD:crud{
		CREATE:   "create",
		RETRIEVE: "retrieve",
		UPDATE:   "update",
		DELETE:   "delete",
	},


}


