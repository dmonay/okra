package handlers

import (
	// "fmt"
	"github.com/dmonay/do-work-api/authentication"
	"github.com/dmonay/do-work-api/common"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func (dw *DoWorkResource) Login(c *gin.Context) {

	var creds common.Credentials

	c.Bind(&creds)

	uname := creds.Username

	// user := common.User{0, uname, creds.Password}
	authentication.CreateCookie(c.Writer)

	c.JSON(200, "You have successfully logged in, "+uname)

	// dw.db.Save(&user)

}
