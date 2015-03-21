package handlers

import (
	// "fmt"
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) GetOneUser(c *gin.Context) {

	gid := c.Params.ByName("gid")

	var result common.UsersObj
	err := dw.mongo.C("Users").Find(bson.M{"gid": gid}).One(&result)
	if err != nil {
		CheckErr(err, "Failed to get user", c)
		return
	}

	c.JSON(200, SuccessMsg{result.Id})

}
