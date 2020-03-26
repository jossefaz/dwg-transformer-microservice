package model

import (
	"encoding/json"
	"fmt"
)

func Retrieve(atts interface{}, db *CDb, keyval map[string]interface{}) ([]byte, error) {
	errors := db.Where(keyval).Find(atts).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(atts)
	return b, nil
}

func Create(atts interface{}, db *CDb) ([]byte, error) {
	errors := db.Create(atts).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(string(db.RowsAffected), " rows were updated")), nil
}
