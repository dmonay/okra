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

	// db.CreateTable(&common.User{})

	// // construct a gorp DbMap
	// UserDb := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// // add a table
	// UserDb.AddTable(common.User{}).SetKeys(true, "Id")

	// // create the table
	// err = UserDb.CreateTablesIfNotExists()
	// checkErr(err, "Create tables failed")

	// test insert
	// testQuery := &common.User{0, "John", "pwd"}
	// err = UserDb.Insert(testQuery)

	return db, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
