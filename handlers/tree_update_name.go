package handlers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) UpdateTreeName(c *gin.Context) {
	org := c.Params.ByName("organization")
	treeId := c.Params.ByName("treeid")
	id := bson.ObjectIdHex(treeId)

	type TreeNameJson struct {
		TreeName string `json:"treename"`
	}

	var reqBody TreeNameJson
	c.Bind(&reqBody)
	newName := reqBody.TreeName

	querier := bson.M{"_id": id}
	updateName := bson.M{"$set": bson.M{"treename": newName}}
	err := dw.mongo.C(org).Update(querier, updateName)
	if err != nil {
		CheckErr(err, "Mongo failed to update tree name", c)
		return
	}

	c.JSON(201, SuccessMsg{"Success"})
}
