package handlers

import (
	// "fmt"
	// "github.com/codegangsta/martini-contrib/binding"
	"github.com/dmonay/do-work-api/common"
	// "github.com/dmonay/do-work-api/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "net/http"
)

func (dw *DoWorkResource) Register(c *gin.Context) {

	var creds common.Credentials

	c.Bind(&creds)

	uname := creds.Username

	user := common.User{0, uname, creds.Password}

	c.JSON(201, "you have registered, "+uname)

	dw.db.Save(&user)
}

type DoWorkResource struct {
	db gorm.DB
}
