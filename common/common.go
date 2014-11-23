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
	DbUser     string
	DbPassword string
	DbName     string
}
