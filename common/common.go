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
}

type OrganizationJson struct {
	Organization string `json:"organization"`
	UserId       string `json:"userId" bson:"_id,omitempty"`
}

type TreeJson struct {
	TreeName  string `json:"treeName"`
	Timeframe string `json:"timeframe"`
	UserId    string `json:"userId" bson:"_id,omitempty"`
}

type MembersJson struct {
	Members []Member `json:"members"`
}

type Member struct {
	Username string `json:"userName"`
	UserId   string `json:"userId"`
	Role     string `json:"role"`
}

type ObjectiveJson struct {
	Id      string                       `json:"id"`
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
