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
