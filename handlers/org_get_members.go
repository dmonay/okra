package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) GetMembers(c *gin.Context) {
	org := c.Params.ByName("org")

	var result common.MembersInOrg
	err := dw.mongo.C(org).Find(bson.M{"name": "membersArray"}).One(&result)
	if err != nil {
		CheckErr(err, "Failed to retrieve members in "+org+" org", c)
		return
	}

	c.JSON(200, SuccessMsg{result.Members})
}
