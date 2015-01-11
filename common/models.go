package common

import (
	"gopkg.in/mgo.v2/bson"
)

type UserTree struct {
	TreeName string
	TreeId   bson.ObjectId
}

type OrgMembers struct {
	Members []Member `bson:"members"`
	Name    string   `bson:"name"`
}

type Member struct {
	Username string `json:"userName"`
	UserId   string `json:"userId",omitempty`
	Role     string `json:"role"`
}

type UsersObj struct {
	Username string        `json:"username"`
	Orgs     []string      `bson:"orgs" omitempty`
	Trees    []string      `bson:"trees" omitempty`
	Id       bson.ObjectId `bson:"_id"`
}

type OkrTree struct {
	Id         bson.ObjectId `bson:"_id"`
	Type       string
	OrgName    string
	Mission    string
	Members    []Member
	Active     bool
	Timeframe  string
	TreeName   string
	Objectives []ObjectiveMongo
}

type ObjectiveMongo struct {
	Name       string
	Body       string
	Completed  bool
	Members    []Member
	KeyResults []KeyResultsModel
}

type KeyResultsModel struct {
	Name      string
	Body      string
	Completed bool
	Members   []Member
	Priority  string
	Tasks     []TasksModel
}

type TasksModel struct {
	Name      string
	Body      string
	Completed bool
	Members   []Member
	Priority  string
}

type TreeInOrg struct {
	Name   string        `bson:"treename"`
	Id     bson.ObjectId `bson:"_id"`
	Active bool          `bson:"active"`
}
