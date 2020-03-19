package model

import (
	"fmt"
	tables "github.com/yossefaz/go_struct"
)


func StatusUpdate(db *CDb, where map[string]interface{}, update map[string]interface{}) ([]byte, error){
	errors := db.Model(tables.Cad_check_status{}).Where(where).Updates(update).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(string(db.RowsAffected), " rows were updated")), nil
}




