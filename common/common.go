package common

import (
	"gopkg.in/mgo.v2/bson"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbName     string
}

type UserJson struct {
	Username string   `json:"username"`
	Orgs     []string `bson:"orgs" omitempty`
	Trees    []string `bson:"trees" omitempty`
}

type MissionJson struct {
	Mission string `json:"mission"`
	TreeId  string `json:"treeId,omitempty"`
}

type OrganizationJson struct {
	Organization string `json:"organization"`
	UserId       string `json:"userId" bson:"_id,omitempty"`
	UserName     string `json:"username"`
	UserRole     string `json:"role"`
}

type TreeJson struct {
	TreeName  string `json:"treeName"`
	Timeframe string `json:"timeframe"`
	UserName  string `json:"username"`
	UserId    string `json:"userId" bson:"_id,omitempty"`
	UserRole  string `json:"role"`
}

type MembersJson struct {
	UpdateTree bool     `json:"updateTree,omitempty"`
	TreeName   string   `json:"treeName,omitempty"`
	TreeId     string   `json:"treeId,omitempty"`
	Members    []Member `json:"members"`
}

type MembersJsonDelete struct {
	UpdateTree bool     `json:"updateTree,omitempty"`
	TreeId     string   `json:"treeId,omitempty"`
	Members    []string `json:"members"`
}

type Member struct {
	Username string `json:"userName"`
	UserId   string `json:"userId",omitempty`
	Role     string `json:"role"`
}

type ObjectiveJson struct {
	Id      string                       `json:"id"`
	TreeId  string                       `json:"treeId,omitempty"`
	Name    string                       `json:"name"`
	Body    string                       `json:"body"`
	Active  bool                         `json:"active"`
	Members map[string]map[string]string `json:"members"`
}

type KeyResultJson struct {
	Id       string                       `json:"id"`
	Members  map[string]map[string]string `json:"members"`
	Name     string                       `json:"name"`
	Body     string                       `json:"body"`
	Active   bool                         `json:"active"`
	Priority string                       `json:"priority"`
}

type UserOrgs struct {
	OrgName string
}

type UserTree struct {
	TreeName string
	TreeId   bson.ObjectId
}

type OrgMembers struct {
	Members []Member `bson:"members"`
	Name    string   `bson:"name"`
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
	Objective1 ObjectiveMongo
	Objective2 ObjectiveMongo
	Objective3 ObjectiveMongo
	Objective4 ObjectiveMongo
	Objective5 ObjectiveMongo
}

type ObjectiveMongo struct {
	Name    string                       `json:"name"`
	Body    string                       `json:"body"`
	Active  bool                         `json:"active"`
	Members map[string]map[string]string `json:"members"`
}

type TreeInOrg struct {
	Name   string        `bson:"treename"`
	Id     bson.ObjectId `bson:"_id"`
	Active bool          `bson:"active"`
}
