package handlers

import (
	"fmt"
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) CreateOrg(c *gin.Context) {

	var reqBody common.OrganizationJson

	c.Bind(&reqBody)

	org := reqBody.Organization
	err := dw.mongo.C(org).Insert(&OkrTree{
		org,
		"",
		"",
		false,
		"",
		Objective{""},
		Objective{""},
		Objective{""},
		Objective{""},
		Objective{""},
	})
	CheckErr(err, "Mongo failed to create collection for "+org+" organization")

	c.JSON(201, "You have successfully created an organization")
}

func (dw *DoWorkResource) CreateTree(c *gin.Context) {

	// not actually creating a tree, as that was created in CreateOrg.
	// Simply updating the 'timeframe' and 'active' fields

	org := c.Params.ByName("organization")
	var reqBody common.TreeJson

	c.Bind(&reqBody)

	timeframe := reqBody.Timeframe

	colQuerier := bson.M{"orgname": org}

	// 1. update the timeframe of the tree
	updateTimeframe := bson.M{"$set": bson.M{"timeframe": timeframe}}
	err := dw.mongo.C(org).Update(colQuerier, updateTimeframe)
	CheckErr(err, "Mongo failed to update timeframe")

	// 2. set status to true
	updateStatus := bson.M{"$set": bson.M{"active": true}}

	err2 := dw.mongo.C(org).Update(colQuerier, updateStatus)
	CheckErr(err2, "Mongo failed to set 'active' status to true")

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

func (dw *DoWorkResource) UpdateMembers(c *gin.Context) {

	org := c.Params.ByName("organization")
	var reqBody common.MembersJson

	c.Bind(&reqBody)

	members := reqBody.Members

	colQuerier := bson.M{"orgname": org}
	addMembers := bson.M{"$set": bson.M{"members": members}}
	err := dw.mongo.C(org).Update(colQuerier, addMembers)
	CheckErr(err, "Mongo failed to add members to "+org+" organization")

	c.JSON(201, "You have successfully added members to the "+org+" organization")
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
	// body := reqBody.Body
	// members := reqBody.Members
	name := obj + "." + id

	colQuerier := bson.M{"orgname": org}
	addKeyResult := bson.M{"$set": bson.M{name: reqBody}}
	err := dw.mongo.C(org).Update(colQuerier, addKeyResult)
	CheckErr(err, "Mongo failed to add key result")

	fmt.Println(reqBody)

	c.JSON(201, "You have successfully added a key result")
}

type OkrTree struct {
	OrgName    string
	Mission    string
	Members    string
	Active     bool
	Timeframe  string
	Objective1 Objective
	Objective2 Objective
	Objective3 Objective
	Objective4 Objective
	Objective5 Objective
}

type Objective struct {
	Name string
}
