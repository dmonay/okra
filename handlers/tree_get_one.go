package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) GetTree(c *gin.Context) {
	org := c.Params.ByName("organization")
	treeId := c.Params.ByName("treeid")
	id := bson.ObjectIdHex(treeId)

	var result common.OkrTree

	err := dw.mongo.C(org).Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		CheckErr(err, "Failed to retrieve tree from Mongo", c)
		return
	}
	result.Id = id

	c.JSON(200, SuccessMsg{result})
}
