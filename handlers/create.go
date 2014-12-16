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
	memArray := []string{}
	membersDoc := &OrgMembers{memArray, "membersArray"}
	err2 := dw.mongo.C(org).Insert(membersDoc)
	CheckErr(err2, "Mongo failed to create collection with the empty members array")

	// 2. Add organization to user's document in Users
	colQuerier := bson.M{"_id": objId}
	updateTimeframe := bson.M{"$push": bson.M{"orgs": org}}
	err := dw.mongo.C("Users").Update(colQuerier, updateTimeframe)
	CheckErr(err, "Mongo failed to add orgnization to user's document in Users")

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3333")
	c.JSON(201, "You have successfully created an organization")
}

type OrgMembers struct {
	Members []string `bson:"members"`
	Name    string   `bson:"name"`
}

// func (dw *DoWorkResource) CreateOrgOpts(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3333")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
// }

func (dw *DoWorkResource) CreateTree(c *gin.Context) {

	var reqBody common.TreeJson
	c.Bind(&reqBody)

	org := c.Params.ByName("organization")
	timeframe := reqBody.Timeframe
	treeName := reqBody.TreeName
	objId := bson.ObjectIdHex(reqBody.UserId)
	members := []string{}

	// 1. Create tree and upsert it into the org collection
	treeStruct := &OkrTree{
		org,
		"",
		members,
		true,
		timeframe,
		treeName,
		Objective{""},
		Objective{""},
		Objective{""},
		Objective{""},
		Objective{""},
	}
	colQuerier := bson.M{"treename": treeName}
	upsertTree := bson.M{"$set": treeStruct}
	info, err := dw.mongo.C(org).Upsert(colQuerier, upsertTree)
	CheckErr(err, "Mongo failed to create collection for "+org+" organization")

	// 2. Update user's doc in Users with the ObjId and name of the tree
	treeId := info.UpsertedId.(bson.ObjectId)
	tree := common.UserTree{treeName, treeId}
	colQuerier2 := bson.M{"_id": objId}
	updateTimeframe := bson.M{"$push": bson.M{"trees": tree}}
	err2 := dw.mongo.C("Users").Update(colQuerier2, updateTimeframe)
	CheckErr(err2, "Mongo failed to add tree to user's document in Users")

	c.JSON(201, "You have successfully created a tree with the "+timeframe+" timeframe for the "+org+" organization")
}

func (dw *DoWorkResource) UpdateMission(c *gin.Context) {

	org := c.Params.ByName("organization")
	var reqBody common.MissionJson

	c.Bind(&reqBody)

	mission := reqBody.Mission

	colQuerier := bson.M{"orgname": org}
	setMission := bson.M{"$set": bson.M{"mission": mission}}
	err := dw.mongo.C(org).Update(colQuerier, setMission)
	CheckErr(err, "Mongo failed to update mission")

	c.JSON(201, "You have successfully added a mission")
}

func (dw *DoWorkResource) AddMembers(c *gin.Context) {
	org := c.Params.ByName("organization")
	var reqBody common.MembersJson
	c.Bind(&reqBody)

	// update tree's members array if tree ID provided
	if reqBody.UpdateTree {
		id := bson.ObjectIdHex(reqBody.Tree)

		colQuerier := bson.M{"_id": id}
		for _, value := range reqBody.Members {
			addMembers := bson.M{"$push": bson.M{"members": value}}
			err := dw.mongo.C(org).Update(colQuerier, addMembers)
			CheckErr(err, "Mongo failed to add members to the provided tree")
		}

		c.JSON(201, "You have successfully added members to the tree")

		// otherwise update org's members array
	} else {
		colQuerier := bson.M{"name": "membersArray"}
		for _, value := range reqBody.Members {
			addMembers := bson.M{"$push": bson.M{"members": value}}
			err := dw.mongo.C(org).Update(colQuerier, addMembers)
			CheckErr(err, "Mongo failed to add members to "+org+" organization")
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

		id := bson.ObjectIdHex(reqBody.Tree)

		colQuerier := bson.M{"_id": id}
		for _, value := range reqBody.Members {
			fmt.Println("user's id: ", value)
			removeMembers := bson.M{"$pull": bson.M{"members": bson.M{"userid": value}}}
			err := dw.mongo.C(org).Update(colQuerier, removeMembers)
			CheckErr(err, "Mongo failed to remove members from the provided tree")
		}
		c.JSON(201, "You have successfully removed members from the tree")

		// otherwise remove member from org's members array
	} else {
		colQuerier := bson.M{"name": "membersArray"}
		for _, value := range reqBody.Members {
			removeMembers := bson.M{"$pull": bson.M{"members": bson.M{"userid": value}}}
			err := dw.mongo.C(org).Update(colQuerier, removeMembers)
			CheckErr(err, "Mongo failed to remove members from "+org+" organization")
		}
		c.JSON(201, "You have successfully removed members from the "+org+" organization")
	}
}

func (dw *DoWorkResource) UpdateObjective(c *gin.Context) {

	org := c.Params.ByName("organization")
	var reqBody common.ObjectiveJson

	c.Bind(&reqBody)

	id := reqBody.Id
	// body := reqBody.Body
	// members := reqBody.Members

	colQuerier := bson.M{"orgname": org}
	addObjective := bson.M{"$set": bson.M{id: reqBody}}
	err := dw.mongo.C(org).Update(colQuerier, addObjective)
	CheckErr(err, "Mongo failed to add objective")

	c.JSON(201, "You have successfully added an objective to the "+org+" organization")
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

type OkrTree struct {
	OrgName    string
	Mission    string
	Members    []string
	Active     bool
	Timeframe  string
	TreeName   string
	Objective1 Objective
	Objective2 Objective
	Objective3 Objective
	Objective4 Objective
	Objective5 Objective
}

type Objective struct {
	Name string
}
