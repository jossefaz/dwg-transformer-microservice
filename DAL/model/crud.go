package model

import "encoding/json"

func Retrieve (atts interface{}, db *CDb, keyval map[string]interface{})([]byte, error) {
	errors := db.Where(keyval).Find(atts).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(atts)
	return b, nil
}
