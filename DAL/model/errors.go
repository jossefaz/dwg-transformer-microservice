package model

import (
	"encoding/json"
)

type Cad_check_errors struct {
	Id int
	check_status int
	error_code Timestamp
}

type LUT_cad_errors struct {
	Id int
	Func_name string
}

func (Cad_check_errors) TableName() string {
	return "CAD_check_errors"
}

func (LUT_cad_errors) TableName() string {
	return "LUT_cad_errors"
}



func ErrorsRetrieve(db *CDb, keyval map[string]interface{}) ([]byte, error){
	atts :=  []Cad_check_errors{}
	errors := db.Where(keyval).Find(&atts).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(atts)
	return b, nil
}



func Lut_Error_Retrieve(db *CDb, keyval map[string]interface{}) (map[string]interface{}){
	atts :=  LUT_cad_errors{}
	copyKeyval := make(map[string]interface{})
	for errorName, _ := range keyval {
		db.Where("func_name = ?", errorName).Find(&atts).GetErrors()
		copyKeyval[errorName] = atts.Id
	}

	return copyKeyval
}

func ErrorsCreate(db *CDb, FkId map[string]interface{}, keyval map[string]interface{}) ([]byte, error){


	//keyval = Lut_Error_Retrieve(db, keyval)
	//for _, errorCode := range keyval {
	//	atts :=  Cad_check_errors{}
	//
	//	atts.Id = FkId["check_status_id"]
	//
	//	errors := db.Create(&atts).GetErrors()
	//
	//}
	//
	//
	//err := HandleDBErrors(errors)
	//if err != nil {
	//	return nil, err
	//}
	//b, _ := json.Marshal(atts)
	return []byte{}, nil
}

