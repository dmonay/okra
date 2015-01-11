package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) UpdateMission(c *gin.Context) {

	org := c.Params.ByName("organization")
	var reqBody common.MissionJson

	c.Bind(&reqBody)

	mission := reqBody.Mission
	treeId := bson.ObjectIdHex(reqBody.TreeId)

	colQuerier := bson.M{"_id": treeId}
	setMission := bson.M{"$set": bson.M{"mission": mission}}
	err := dw.mongo.C(org).Update(colQuerier, setMission)
	CheckErr(err, "Mongo failed to update mission")

	c.JSON(201, mission)
}
