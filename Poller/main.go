package Poller

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:Dev123456!@(localhost)/dbname?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		fmt.Println("Cannot connect to DB")
	}
	defer db.Close()
}