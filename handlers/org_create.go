package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) CreateOrg(c *gin.Context) {

	var reqBody common.OrganizationJson

	c.Bind(&reqBody)

	org := reqBody.Organization
	objId := bson.ObjectIdHex(reqBody.UserId)

	// 1. Create collection with members document
	memObj := common.Member{reqBody.UserName, reqBody.UserId, "admin"}
	arrayOfMembers := []common.Member{memObj}

	membersDoc := &common.OrgMembers{arrayOfMembers, "membersArray"}
	err2 := dw.mongo.C(org).Insert(membersDoc)
	if err2 != nil {
		CheckErr(err2, "Mongo failed to create collection with the empty members array", c)
		return
	}

	// 2. Add organization to user's document in Users
	colQuerier := bson.M{"_id": objId}
	updateTimeframe := bson.M{"$push": bson.M{"orgs": org}}
	err := dw.mongo.C("Users").Update(colQuerier, updateTimeframe)
	if err != nil {
		CheckErr(err, "Mongo failed to add organization to user's document in Users", c)
		return
	}

	c.JSON(201, SuccessMsg{"You have successfully created an organization"})
}
