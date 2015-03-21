package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) Register(c *gin.Context) {

	var user common.UserJson
	c.Bind(&user)

	uname := user.Username
	displayName := user.DisplayName
	gid := user.Gid
	orgs := []string{}
	trees := []string{}
	id := bson.NewObjectId()

	err := dw.mongo.C("Users").Insert(&common.UsersObj{uname, orgs, trees, displayName, gid, id})
	if err != nil {
		CheckErr(err, "User not added to Users collection", c)
		return
	}

	c.JSON(201, SuccessMsg{id})
}
