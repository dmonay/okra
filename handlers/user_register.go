package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
)

func (dw *DoWorkResource) Register(c *gin.Context) {

	type User struct {
		Username string
	}

	var user common.UserJson

	c.Bind(&user)

	uname := user.Username
	orgs := []string{}
	trees := []string{}

	err := dw.mongo.C("Users").Insert(&common.UserJson{uname, orgs, trees})
	if err != nil {
		CheckErr(err, "User not added to Users collection", c)
		return
	}

	c.JSON(201, SuccessMsg{"You have registered, " + uname})
}
