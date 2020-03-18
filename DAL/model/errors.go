package model

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type CAD_check_errors struct {
	Check_status_id int
	Error_code int
}

type LUT_cad_errors struct {
	Id int
	Func_name string
}

func (CAD_check_errors) TableName() string {
	return "CAD_check_errors"
}

func (LUT_cad_errors) TableName() string {
	return "LUT_cad_errors"
}



func ErrorsRetrieve(db *CDb, keyval map[string]interface{}) ([]byte, error){
	atts :=  []CAD_check_errors{}
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
	for errorName, errorval := range keyval {
		testval := parsInt(errorval)
		if testval == 1 {
			db.Where("func_name = ?", errorName).Find(&atts).GetErrors()
			copyKeyval[errorName] = atts.Id
		}
	}
	return copyKeyval
}

func parsInt(val interface{}) int {
	var testval int
	if reflect.TypeOf(val).Kind() == reflect.Float64 {
		testval =  int(val.(float64))
	}
	if reflect.TypeOf(val).Kind() == reflect.Int {
		testval =  val.(int)
	}
	return testval
}

func ErrorsCreate(db *CDb, FkId map[string]interface{}, keyval map[string]interface{}) ([]byte, error){
	keyval = Lut_Error_Retrieve(db, keyval)
	for _, errorCode := range keyval {
		atts :=  CAD_check_errors{}
		checkId := parsInt(FkId["check_status_id"])
		atts.Check_status_id = checkId
		testval := parsInt(errorCode)
		atts.Error_code = testval
		rows, _ := Create(atts, db)
		fmt.Println(rows)
	}
	return []byte{}, nil
}

