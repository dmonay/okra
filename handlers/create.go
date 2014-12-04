package handlers

import (
	// "fmt"
	// "github.com/dmonay/do-work-api/authentication"
	"github.com/dmonay/do-work-api/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) CreateOrg(c *gin.Context) {

	var reqBody common.OrganizationJson

	c.Bind(&reqBody)

	org := reqBody.Organization
	members := make([]string, 10)
	err := dw.mongo.C(org).Insert(&OkrTree{
		org,
		"",
		members,
		false,
		"",
		Objective{""},
		Objective{""},
		Objective{""},
		Objective{""},
		Objective{""},
	})
	if err != nil {
		panic(err)
	}

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
	if err != nil {
		panic(err)
	}

	// 2. set status to true
	updateStatus := bson.M{"$set": bson.M{"active": true}}

	err2 := dw.mongo.C(org).Update(colQuerier, updateStatus)
	if err2 != nil {
		panic(err)
	}

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
	if err != nil {
		panic(err)
	}

	c.JSON(201, "You have successfully added a mission")
}

type OkrTree struct {
	OrgName    string
	Mission    string
	Members    []string
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
