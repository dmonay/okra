package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) DeleteMembers(c *gin.Context) {
	org := c.Params.ByName("organization")
	var reqBody common.MembersJsonDelete
	c.Bind(&reqBody)

	// remove member from tree's members array if tree ID provided
	if reqBody.UpdateTree {

		id := bson.ObjectIdHex(reqBody.TreeId)

		colQuerier := bson.M{"_id": id}
		for _, value := range reqBody.Members {
			removeMembers := bson.M{"$pull": bson.M{"members": bson.M{"userid": value}}}
			err := dw.mongo.C(org).Update(colQuerier, removeMembers)
			CheckErr(err, "Mongo failed to remove members from the provided tree")

			// Remove trees from user's doc in Users
			memberId := bson.ObjectIdHex(value)
			treeId := bson.ObjectIdHex(reqBody.TreeId)
			colQuerier2 := bson.M{"_id": memberId}
			updateUsersDoc := bson.M{"$pull": bson.M{"trees": bson.M{"treeid": treeId}}}
			err2 := dw.mongo.C("Users").Update(colQuerier2, updateUsersDoc)
			CheckErr(err2, "Mongo failed to remove tree from user's document in Users")
		}
		c.JSON(200, "You have successfully removed members from the tree")

		// otherwise remove member from org's members array
	} else {
		colQuerier := bson.M{"name": "membersArray"}
		for _, value := range reqBody.Members {
			removeMembers := bson.M{"$pull": bson.M{"members": bson.M{"userid": value}}}
			err := dw.mongo.C(org).Update(colQuerier, removeMembers)
			CheckErr(err, "Mongo failed to remove members from "+org+" organization")

			// Remove org from user's doc in Users
			memberId := bson.ObjectIdHex(value)
			colQuerier2 := bson.M{"_id": memberId}
			updateUsersDoc := bson.M{"$pull": bson.M{"orgs": org}}
			err2 := dw.mongo.C("Users").Update(colQuerier2, updateUsersDoc)
			CheckErr(err2, "Mongo failed to remove org from user's document in Users")
		}
		c.JSON(200, "You have successfully removed members from the "+org+" organization")
	}
}
