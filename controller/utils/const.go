package utils

type channels struct {
	CheckDWG string
	CheckedDWG string
	ConvertDWG string
	ConvertedDWG string
	Dal_Req string
	Dal_Res string
}

type headers map[string]map[string]interface{}

type constants struct {
	CheckDWG string
	From string
	Channels channels
	Headers headers
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

}


