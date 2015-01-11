package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) GetAllOrgs(c *gin.Context) {
	org := c.Params.ByName("userid")
	id := bson.ObjectIdHex(org)

	var result common.UsersObj
	err4 := dw.mongo.C("Users").Find(bson.M{"_id": id}).One(&result)
	CheckErr(err4, "Failed to retrieve orgs from Users coll in Mongo")

	c.JSON(200, result.Orgs)
}
