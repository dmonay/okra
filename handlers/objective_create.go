package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) CreateObjective(c *gin.Context) {

	org := c.Params.ByName("organization")
	var reqBody common.ObjectiveJson
	c.Bind(&reqBody)

	arrayOfMembers := []common.Member{}
	for _, value := range reqBody.Members {
		// Get userId of user
		var result common.UsersObj
		err := dw.mongo.C("Users").Find(bson.M{"username": value.Username}).One(&result)
		if err != nil {
			CheckErr(err, "Mongo failed to find the "+value.Username+"'s doc in Users", c)
			return
		}

		memObj := common.Member{value.Username, result.Id.Hex(), value.Role}
		arrayOfMembers = append(arrayOfMembers, memObj)
	}

	treeId := bson.ObjectIdHex(reqBody.TreeId)
	obj := common.ObjectiveMongo{
		Id:        reqBody.Id,
		Name:      reqBody.Name,
		Body:      reqBody.Body,
		Completed: reqBody.Completed,
		Members:   arrayOfMembers,
	}

	colQuerier := bson.M{"_id": treeId}
	addObjective := bson.M{"$push": bson.M{"objectives": obj}}
	err2 := dw.mongo.C(org).Update(colQuerier, addObjective)
	if err2 != nil {
		CheckErr(err2, "Mongo failed to add objective", c)
		return
	}

	c.JSON(201, SuccessMsg{"You have successfully added an objective!"})
}
