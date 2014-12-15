package common

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbName     string
}

type UserJson struct {
	Username string `json:"username"`
}

type MissionJson struct {
	Mission string `json:"mission"`
}

type OrganizationJson struct {
	Organization string `json:"organization"`
	UserId       string `json:"userId" bson:"_id,omitempty"`
}

type TreeJson struct {
	Timeframe string `json:"timeframe"`
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

type UserOrganizations struct {
	UserOrg []UserOrgs
}

type UserOrgs struct {
	OrgName string
	OrgId   bson.ObjectId
}
