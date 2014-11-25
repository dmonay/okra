package handlers

import (
	// "fmt"
	"github.com/dmonay/do-work-api/common"
	"github.com/gin-gonic/gin"
)

func (dw *DoWorkResource) Logout(c *gin.Context) {

	var creds common.Credentials

	c.Bind(&creds)

	// uname := creds.Username

	// user := common.User{0, uname, creds.Password}

	c.JSON(200, "You have successfully logged out")

	// dw.db.Save(&user)
}
