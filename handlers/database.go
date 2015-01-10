package handlers

import (
	"gopkg.in/mgo.v2"
	"log"
)

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
