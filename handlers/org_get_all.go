package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) GetAllOrgs(c *gin.Context) {
	userid := c.Params.ByName("userid")
	id := bson.ObjectIdHex(userid)

	var result common.UsersObj
	err := dw.mongo.C("Users").Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		CheckErr(err, "Failed to retrieve orgs from Users coll in Mongo", c)
		return
	}

	c.JSON(200, SuccessMsg{result.Orgs})
}
