package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) AddMembers(c *gin.Context) {
	org := c.Params.ByName("organization")
	var reqBody common.MembersJson
	c.Bind(&reqBody)

	// update tree's members array if tree ID provided
	if reqBody.UpdateTree {
		treeId := bson.ObjectIdHex(reqBody.TreeId)

		colQuerier := bson.M{"_id": treeId}
		for _, value := range reqBody.Members {

			// Get userId of user
			var result common.UsersObj
			err := dw.mongo.C("Users").Find(bson.M{"username": value.Username}).One(&result)
			if err != nil {
				CheckErr3(err, "Mongo failed to find "+value.Username+"'s doc in Users", c)
				return
			}

			value.UserId = result.Id.Hex()

			// Add member to tree
			addMembers := bson.M{"$push": bson.M{"members": value}}
			err2 := dw.mongo.C(org).Update(colQuerier, addMembers)
			if err2 != nil {
				CheckErr3(err, "Mongo failed to add members to the provided tree", c)
				return
			}

			// Update user's doc in Users with the ObjId and name of the tree
			tree := common.UserTree{reqBody.TreeName, treeId}
			colQuerier2 := bson.M{"_id": result.Id}
			updateUsersDoc := bson.M{"$push": bson.M{"trees": tree}}
			err3 := dw.mongo.C("Users").Update(colQuerier2, updateUsersDoc)
			if err3 != nil {
				CheckErr3(err3, "Mongo failed to add tree to user's document in Users", c)
				return
			}

		}

		c.JSON(201, SuccessMsg{"You have successfully added members to the tree"})

		// otherwise update org's members array
	} else {
		colQuerier := bson.M{"name": "membersArray"}
		for _, value := range reqBody.Members {

			// Get userId of user
			var result common.UsersObj
			err := dw.mongo.C("Users").Find(bson.M{"username": value.Username}).One(&result)
			if err != nil {
				CheckErr3(err, "Mongo failed to find "+value.Username+"'s doc in Users", c)
				return
			}
			value.UserId = result.Id.Hex()

			// add user to the org
			addMembers := bson.M{"$push": bson.M{"members": value}}
			err3 := dw.mongo.C(org).Update(colQuerier, addMembers)
			if err3 != nil {
				CheckErr3(err3, "Mongo failed to add members to "+org+" organization", c)
				return
			}

			// Update user's doc in Users with the name of the org
			colQuerier2 := bson.M{"username": value.Username}
			updateUsersDoc := bson.M{"$push": bson.M{"orgs": org}}
			err2 := dw.mongo.C("Users").Update(colQuerier2, updateUsersDoc)
			if err2 != nil {
				CheckErr3(err2, "Mongo failed to add org to user's document in Users", c)
				return
			}
		}

		c.JSON(201, SuccessMsg{"You have successfully added members to the " + org + " organization"})
	}
}
