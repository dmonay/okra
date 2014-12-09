package handlers

import (
	// "database/sql"
	// "fmt"
	"github.com/dmonay/okra/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"log"
)

func InitSqlDb(cfg common.Config) (gorm.DB, error) {

	connection := cfg.DbUser + ":" + cfg.DbPassword + "@/" + cfg.DbName

	// connect to db
	db, err := gorm.Open("mysql", connection)
	CheckErr(err, "gorm.Open failed")

	return db, nil
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func InitMongo() (*mgo.Database, error) {
	session, err := mgo.Dial("localhost:27017")
	CheckErr(err, "mongo connection failed")

	db := session.DB("testing")
	CheckErr(err, "mongo opening database failed")

	return db, nil
}
