package model

import (
	"encoding/json"
	"reflect"
	tables "github.com/yossefaz/go_struct"
)

func ErrorsRetrieve(db *CDb, keyval map[string]interface{}) ([]byte, error){
	atts :=  []tables.CAD_check_errors{}
	errors := db.Where(keyval).Find(&atts).GetErrors()
	err := HandleDBErrors(errors)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(atts)
	return b, nil
}

func Lut_Error_Retrieve(db *CDb, keyval map[string]interface{}) (map[string]interface{}){
	atts :=  tables.LUT_cad_errors{}
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

func checkIfExist(db *CDb, id int, errorCode int) bool {
	atts :=  tables.CAD_check_errors{}
	if db.Where(&tables.CAD_check_errors{Check_status_id: id, Error_code: errorCode}).First(&atts).RecordNotFound() {
		return false
	}
	return true
}

func ErrorsCreate(db *CDb, FkId map[string]interface{}, keyval map[string]interface{}) ([]byte, error){
	keyval = Lut_Error_Retrieve(db, keyval)
	for _, errorCode := range keyval {
		checkId := parsInt(FkId["check_status_id"])
		errVal := parsInt(errorCode)
		if !checkIfExist(db, checkId, errVal) {
			atts :=  tables.CAD_check_errors{}
			atts.Check_status_id = checkId
			atts.Error_code = errVal
			_, err := Create(atts, db)
			HandleDBErrors([]error{err})
		}
	}
	return []byte{}, nil
}

