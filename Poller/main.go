package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
type Timestamp time.Time

type Attachements struct {
	reference int
	status int
	statusDate Timestamp
	path string
}

func main() {
	db, err := gorm.Open("mysql", "root:Dev123456!@(localhost)/dwg_transformer?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		fmt.Println("Cannot connect to DB", err)
		panic("Cannot connect to DB")
	}
	db.DB()
	db.DB().Ping()
	defer db.Close()
	check:= db.HasTable("Attachements")
	fmt.Println(check)
	//tables := []string{}
	//db.Select(&tables, "SELECT * FROM Attachements")
	//fmt.Println(tables)

	//att := Attachements{}
	//db.First(&att)
	//
	//fmt.Println(att)
	//fmt.Println("------------------------------")
	// get all records
	//db.AutoMigrate(&Attachements{})
	//atts := []Attachements{} // a slice
	//db.Where(&Attachements{status: 0}).Find(&atts)
	//for _, v := range atts {
	//	fmt.Println("reference : ", v.reference)
	//	fmt.Println("path : ", v.path)
	//}

	//dbConn.Find(&activities)
	//fmt.Println(activities)

	//rows, err := db.Model(&Attachements{}).Rows()
	//defer rows.Close()
	//if err != nil {
	//	panic(err)
	//}
	//for rows.Next() {
	//	db.ScanRows(rows, &att)
	//	fmt.Println(att)
	//}


}