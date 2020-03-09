package model

import (
	"encoding/json"
	"fmt"
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

func Att_Retrieve(db *CDb, keyval map[string]interface{}) ([]byte, error){
	atts :=  []Attachements{}
	errors := db.Where(keyval).Find(&atts).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(atts)
	return b, nil
}

func Att_Update(db *CDb, where map[string]interface{}, update map[string]interface{}) ([]byte, error){
	errors := db.Model(Attachements{}).Where(where).Updates(update).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(string(db.RowsAffected), " rows were updated")), nil
}