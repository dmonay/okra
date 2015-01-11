package handlers

import (
	"fmt"
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	// "gopkg.in/mgo.v2"
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
	CheckErr(err2, "Mongo failed to create collection with the empty members array")

	// 2. Add organization to user's document in Users
	colQuerier := bson.M{"_id": objId}
	updateTimeframe := bson.M{"$push": bson.M{"orgs": org}}
	err := dw.mongo.C("Users").Update(colQuerier, updateTimeframe)
	CheckErr(err, "Mongo failed to add organization to user's document in Users")

	c.JSON(201, "You have successfully created an organization")
}

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
	CheckErr(err, "Mongo failed to create tree for "+org+" organization")

	// 2. Update user's doc in Users with the ObjId and name of the tree
	treeId := info.UpsertedId.(bson.ObjectId)
	tree := common.UserTree{treeName, treeId}
	colQuerier2 := bson.M{"_id": objId}
	updateTimeframe := bson.M{"$push": bson.M{"trees": tree}}
	err2 := dw.mongo.C("Users").Update(colQuerier2, updateTimeframe)
	CheckErr(err2, "Mongo failed to add tree to user's document in Users")

	c.JSON(201, treeStruct)
}

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
			CheckErr(err, "Mongo failed to find the "+value.Username+"'s doc in Users")
			value.UserId = result.Id.Hex()

			// Add member to tree
			addMembers := bson.M{"$push": bson.M{"members": value}}
			err2 := dw.mongo.C(org).Update(colQuerier, addMembers)
			CheckErr(err2, "Mongo failed to add members to the provided tree")

			// Update user's doc in Users with the ObjId and name of the tree
			tree := common.UserTree{reqBody.TreeName, treeId}
			colQuerier2 := bson.M{"_id": result.Id}
			updateUsersDoc := bson.M{"$push": bson.M{"trees": tree}}
			err3 := dw.mongo.C("Users").Update(colQuerier2, updateUsersDoc)
			CheckErr(err3, "Mongo failed to add tree to user's document in Users")
		}

		c.JSON(201, "You have successfully added members to the tree")

		// otherwise update org's members array
	} else {
		colQuerier := bson.M{"name": "membersArray"}
		for _, value := range reqBody.Members {

			// Get userId of user
			var result common.UsersObj
			err := dw.mongo.C("Users").Find(bson.M{"username": value.Username}).One(&result)
			CheckErr(err, "Mongo failed to find the "+value.Username+"'s doc in Users")
			value.UserId = result.Id.Hex()

			// add user to the org
			addMembers := bson.M{"$push": bson.M{"members": value}}
			err3 := dw.mongo.C(org).Update(colQuerier, addMembers)
			CheckErr(err3, "Mongo failed to add members to "+org+" organization")

			// Update user's doc in Users with the name of the org
			colQuerier2 := bson.M{"username": value.Username}
			updateUsersDoc := bson.M{"$push": bson.M{"orgs": org}}
			err2 := dw.mongo.C("Users").Update(colQuerier2, updateUsersDoc)
			CheckErr(err2, "Mongo failed to add org to user's document in Users")
		}
		c.JSON(201, "You have successfully added members to the "+org+" organization")
	}
}

func (dw *DoWorkResource) DeleteMembers(c *gin.Context) {
	org := c.Params.ByName("organization")
	var reqBody common.MembersJsonDelete
	c.Bind(&reqBody)

	// remove member from tree's members array if tree ID provided
	if reqBody.UpdateTree {

		id := bson.ObjectIdHex(reqBody.TreeId)

		colQuerier := bson.M{"_id": id}
		for _, value := range reqBody.Members {
			fmt.Println("user's id: ", value)
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

func (dw *DoWorkResource) UpdateObjective(c *gin.Context) {

	org := c.Params.ByName("organization")
	var reqBody common.ObjectiveJson
	c.Bind(&reqBody)

	arrayOfMembers := []common.Member{}
	for _, value := range reqBody.Members {
		// Get userId of user
		var result common.UsersObj
		err := dw.mongo.C("Users").Find(bson.M{"username": value.Username}).One(&result)
		CheckErr(err, "Mongo failed to find the "+value.Username+"'s doc in Users")

		memObj := common.Member{value.Username, result.Id.Hex(), value.Role}
		arrayOfMembers = append(arrayOfMembers, memObj)
	}

	treeId := bson.ObjectIdHex(reqBody.TreeId)
	obj := common.ObjectiveMongo{
		Name:      reqBody.Name,
		Body:      reqBody.Body,
		Completed: reqBody.Completed,
		Members:   arrayOfMembers,
	}

	colQuerier := bson.M{"_id": treeId}
	addObjective := bson.M{"$push": bson.M{"objectives": obj}}
	err := dw.mongo.C(org).Update(colQuerier, addObjective)
	CheckErr(err, "Mongo failed to add objective")

	c.JSON(201, "You have successfully added an objective!")
}

func (dw *DoWorkResource) CreateKeyResult(c *gin.Context) {

	org := c.Params.ByName("organization")
	obj := c.Params.ByName("objective")
	var reqBody common.KeyResultJson

	c.Bind(&reqBody)

	id := reqBody.Id
	resultName := obj + "." + id

	colQuerier := bson.M{"orgname": org}
	addKeyResult := bson.M{"$set": bson.M{resultName: reqBody}}
	err := dw.mongo.C(org).Update(colQuerier, addKeyResult)
	CheckErr(err, "Mongo failed to add key result")

	c.JSON(201, "You have successfully added a key result to "+obj)
}

func (dw *DoWorkResource) GetTrees(c *gin.Context) {
	org := c.Params.ByName("organization")
	treeId := c.Params.ByName("treeid")
	id := bson.ObjectIdHex(treeId)

	var result common.OkrTree

	err := dw.mongo.C(org).Find(bson.M{"_id": id}).One(&result)
	result.Id = id
	CheckErr(err, "Failed to retrieve tree from Mongo")

	c.JSON(200, result)
}

func (dw *DoWorkResource) GetAllTrees(c *gin.Context) {
	org := c.Params.ByName("organization")

	var intermResult []common.OkrTree
	err4 := dw.mongo.C(org).Find(bson.M{"type": "tree"}).All(&intermResult)
	CheckErr(err4, "Failed to retrieve trees in organization "+org+" from Mongo")
	length := len(intermResult)

	result := make([]common.TreeInOrg, length)

	for key, value := range intermResult {
		treeName := value.TreeName
		treeId := value.Id
		active := value.Active

		result[key].Name = treeName
		result[key].Id = treeId
		result[key].Active = active
	}

	c.JSON(200, result)
}
