package database

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *gorp.DbMap {

	// set user and password for your SQL db here
	user := "root"
	pwd := "2Un62VIK"

	// connect to db
	db, err := sql.Open("mysql", user+":"+pwd+"@/do_work")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table
	dbmap.AddTableWithName(User{}).SetKeys(true, "Id")

	// create the table
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	// test insert
	testQuery := &Queries{0, -210.5, 1133.5, -3.5, 42.7, "miles", 19700.4}
	err = dbmap.Insert(testQuery)

	return dbmap
}

type User struct {
	Id       int64
	Username string
	Password string
}
