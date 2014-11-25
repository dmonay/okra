package handlers

import (
	// "database/sql"
	// "fmt"
	"github.com/dmonay/do-work-api/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func InitDb(cfg common.Config) (gorm.DB, error) {

	connection := cfg.DbUser + ":" + cfg.DbPassword + "@/" + cfg.DbName

	// connect to db
	db, err := gorm.Open("mysql", connection)
	checkErr(err, "gorm.Open failed")

	return db, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
