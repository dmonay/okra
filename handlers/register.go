package handlers

import (
	// "fmt"
	"github.com/dmonay/do-work-api/common"
	// "github.com/dmonay/do-work-api/database"
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

func (dw *DoWorkResource) Register(c *gin.Context) {

	var creds common.Credentials

	c.Bind(&creds)

	uname := creds.Username

	user := common.User{0, uname, creds.Password}

	c.JSON(201, "You have registered, "+uname)

	dw.db.Save(&user)

	dw.mongo.C("testData").Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
}

type DoWorkResource struct {
	db    gorm.DB
	mongo *mgo.Database
}

type Person struct {
	Name  string
	Phone string
}
