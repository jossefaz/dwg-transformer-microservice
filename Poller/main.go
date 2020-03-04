package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
type Timestamp time.Time

type Attachements struct {
	Reference int
	Status int
	StatusDate Timestamp
	Path string
}
func (Attachements) TableName() string {
	return "Attachements"
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

	//fmt.Println(att)
	//fmt.Println("------------------------------")
	// get all records
	atts := []Attachements{} // a slice

	db.Where("status = ?", "0").Find(&atts)
	for _, v := range atts {
		fmt.Println("reference : ", v.Reference)
		fmt.Println("path : ", v.Path)
	}
	db.GetErrors()

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