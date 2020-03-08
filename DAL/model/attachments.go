package model

import (
	"encoding/json"
)

type attachements struct {
	Reference int
	Status int
	StatusDate Timestamp
	Path string
}

func (attachements) TableName() string {
	return "Attachements"
}

func Att_Retrieve(db *CDb, keyval map[string]interface{}) []byte{
	atts :=  []attachements{}
	db.Where(keyval).Find(&atts)
	db.GetErrors()
	b, _ := json.Marshal(atts)
	return b
}
