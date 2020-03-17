package model

import (
	"fmt"
)

type Cad_check_status struct {
	Id int
	Status_code int
	Last_update Timestamp
	Path string
	Ref_num int
	System_code int
}


func (Cad_check_status) TableName() string {
	return "CAD_check_status"
}




func StatusUpdate(db *CDb, where map[string]interface{}, update map[string]interface{}) ([]byte, error){
	errors := db.Model(Cad_check_status{}).Where(where).Updates(update).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(string(db.RowsAffected), " rows were updated")), nil
}




