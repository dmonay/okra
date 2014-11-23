package database

import (
	"database/sql"
	// "fmt"
	"github.com/coopernurse/gorp"
	"github.com/dmonay/do-work-api/common"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDb(cfg common.Config) (*gorp.DbMap, error) {

	connection := cfg.DbUser + ":" + cfg.DbPassword + "@/" + cfg.DbName

	// connect to db
	db, err := sql.Open("mysql", connection)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	UserDb := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table
	UserDb.AddTable(common.User{}).SetKeys(true, "Id")

	// create the table
	err = UserDb.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	// test insert
	testQuery := &common.User{0, "John", "pwd"}
	err = UserDb.Insert(testQuery)

	return UserDb, nil
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
