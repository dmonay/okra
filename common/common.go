package common

import (
// "github.com/coopernurse/gorp"
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

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MissionJson struct {
	Mission string `json:"mission"`
}

type OrganizationJson struct {
	Organization string `json:"organization"`
}

type TreeJson struct {
	Timeframe string `json:"timeframe"`
}

type MembersJson struct {
	Members map[string]map[string]string `json:"members"`
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
