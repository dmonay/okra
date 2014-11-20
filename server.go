package main

import (
	// "fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/dmonay/do-work-api/handlers"
)

func main() {
	m := martini.Classic()

	//define the endpoints
	m.Post("/register", binding.Json(handlers.Credentials{}), handlers.Register)
	m.Post("/login", binding.Json(handlers.Credentials{}), handlers.Login)
	m.Post("/logout", binding.Json(handlers.Credentials{}), handlers.Logout)

	m.Run()
}
