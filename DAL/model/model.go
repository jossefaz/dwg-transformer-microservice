package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"os"
	"time"
)

type CDb struct {
	*gorm.DB
}

type Schema struct {
	ConnString string
	Name string
	Dialect string
}

type Timestamp time.Time

func ConnectToDb(dialect string, connString string) *CDb{
	db, err := gorm.Open(dialect, connString)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	db.DB()
	db.DB().Ping()
	var dup = CDb{ db}
	return &dup
}

func (db *CDb) Retrieve( dbQ *globalUtils.DbQuery ) []byte{
	switch dbQ.Table {
	case "Attachments":
		res := Att_Retrieve(db, dbQ.ORMKeyVal)
		return res
	}
	return []byte{}
}
func (db *CDb) Update( dbQ *globalUtils.DbQuery ) {
	switch dbQ.Table {
	case "Attachments":
		Att_Retrieve(db, dbQ.ORMKeyVal)
	}
}
