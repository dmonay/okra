package database

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDb(cfg Config) (*gorp.DbMap, error) {

	connection := cfg.DbUser + ":" + cfg.DbPassword + "@/" + cfg.DbName
	fmt.Println("connection: ", connection)

	// connect to db
	db, err := sql.Open("mysql", connection)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table
	dbmap.AddTable(User{}).SetKeys(true, "Id")

	// create the table
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	// test insert
	testQuery := &User{0, "John", "pwd"}
	err = dbmap.Insert(testQuery)

	return dbmap, nil
}

type User struct {
	Id       int64
	Username string
	Password string
}

type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
