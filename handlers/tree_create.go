package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) CreateTree(c *gin.Context) {

	var reqBody common.TreeJson
	c.Bind(&reqBody)

	org := c.Params.ByName("organization")
	timeframe := reqBody.Timeframe
	treeName := reqBody.TreeName
	objId := bson.ObjectIdHex(reqBody.UserId)
	id := bson.NewObjectId()

	// add member who created Tree to Tree
	memObj := common.Member{reqBody.UserName, reqBody.UserId, "admin"}
	var members []common.Member
	members = append(members, memObj)
	objectives := []common.ObjectiveMongo{}

	// 1. Create tree and upsert it into the org collection
	treeStruct := &common.OkrTree{
		id,
		"tree",
		org,
		"",
		members,
		true,
		timeframe,
		treeName,
		objectives,
	}
	colQuerier := bson.M{"treename": treeName}
	upsertTree := bson.M{"$set": treeStruct}
	info, err := dw.mongo.C(org).Upsert(colQuerier, upsertTree)
	if err != nil {
		CheckErr(err, "Mongo failed to create tree for "+org+" organization", c)
		return
	}

	// 2. Update user's doc in Users with the ObjId and name of the tree
	treeId := info.UpsertedId.(bson.ObjectId)
	tree := common.UserTree{treeName, treeId}
	colQuerier2 := bson.M{"_id": objId}
	updateTimeframe := bson.M{"$push": bson.M{"trees": tree}}
	err2 := dw.mongo.C("Users").Update(colQuerier2, updateTimeframe)
	if err != nil {
		CheckErr(err2, "Mongo failed to add tree to user's document in Users", c)
		return
	}

	result := &common.TreeInOrg{
		treeName,
		id,
		true,
	}
	c.JSON(201, SuccessMsg{result})
}
