package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Attachements struct {
	Reference int
	Status int
	StatusDate Timestamp
	Path string
}

func (Attachements) TableName() string {
	return "Attachements"
}

func Att_Retrieve(db *CDb, keyval map[string]interface{}) []byte{
	atts :=  []Attachements{}
	db.Where(keyval).Find(&atts)
	db.GetErrors()
	b, _ := json.Marshal(atts)
	return b
}

func Att_Update(db *CDb, where map[string]interface{}, update map[string]interface{}) []byte{
	errors := db.Model(Attachements{}).Where(where).Updates(update).GetErrors()
	if len(errors) > 0 {
		var b1 strings.Builder
		for _, err := range errors {
			b1.WriteString(fmt.Sprintln(err))
		}
		return []byte(b1.String())
	}
	return []byte(fmt.Sprintf(string(db.RowsAffected), " rows were updated"))
}