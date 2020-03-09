package model

import (
	"github.com/yossefazoulay/go_utils/test"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"testing"
)

var datbaseQuery  = globalUtils.DbQuery {
	Id: map[string]interface{}{
		"reference" : 5,
	},
	Table :  "Attachments",
	ORMKeyVal: map[string]interface{}{
			"status" : 1,
	},
}



func TestConnectToDb(t *testing.T) {
	_, err := ConnectToDb("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
}

func TestCDb_Retrieve(t *testing.T) {
	cdb, err:= ConnectToDb("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
	_, err1 := cdb.Retrieve(&datbaseQuery)
	test.Ok(t, err1)
}

func TestCDb_Update(t *testing.T) {
	cdb, err:= ConnectToDb("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	test.Ok(t, err)
	_, err1 := cdb.Update(&datbaseQuery)
	test.Ok(t, err1)
}