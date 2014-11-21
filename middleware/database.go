package middleware

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	dbmap.AddTable(User{}).SetKeys(true, "Id")

	// create the table
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	fmt.Println("I ran")
	// test insert
	testQuery := &User{0, "John", "pwd"}
	err = dbmap.Insert(testQuery)

	return dbmap
}

type User struct {
	Id       int64
	Username string
	Password string
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
