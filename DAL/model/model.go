package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	globalUtils "github.com/yossefazoulay/go_utils/utils"
	"strings"
	"time"
	_ "github.com/go-sql-driver/mysql"

)

type CDb struct {
	*gorm.DB
}

type Schema struct {
	ConnString string
	Name string
	Dialect string
}

type DBRes struct {
	modelRes interface{}
	ResType string
}

type Timestamp time.Time

func HandleDBErrors(errs []error) error {
	if len(errs) > 0 {
		var b1 strings.Builder
		for _, err := range errs {
			b1.WriteString(fmt.Sprintln(err))
		}
		return errors.New(b1.String())
	}
	return nil
}

func ConnectToDb(dialect string, connString string) (*CDb, error){
	db, err := gorm.Open(dialect, connString)
	if err != nil {
		return &CDb{}, err
	}
	db.DB()
	db.DB().Ping()
	var dup = CDb{ db}
	return &dup, nil
}

func (db *CDb) Retrieve( dbQ *globalUtils.DbQuery ) ([]byte, error){
	switch dbQ.Table {
	case "Attachments":
		res, err := Att_Retrieve(db, dbQ.ORMKeyVal)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return []byte{}, errors.New("any tables allowed correspond to the requested table name")
}
func (db *CDb) Update( dbQ *globalUtils.DbQuery ) ([]byte, error) {
	switch dbQ.Table {
	case "Attachments":
		res, err := Att_Update(db, dbQ.Id, dbQ.ORMKeyVal)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return []byte{}, errors.New("any tables allowed correspond to the requested table name")
}
