package utils

type channels struct {
	CheckDWG string
	CheckedDWG string
	ConvertDWG string
	ConvertedDWG string
	Dal_Req string
	Dal_Res string
}

type constants struct {
	CheckDWG string
	From string
	Channels channels
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
}
