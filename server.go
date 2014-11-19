package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/dmonay/do-work-api/handlers"
	"github.com/dmonay/worth_my_salt/response"
)

func main() {
	m := martini.Classic()

	//define the endpoints
	m.Post("/register", binding.Json(response.Attribute{}), handlers.Register)
	m.Post("/login", binding.Json(response.Attribute{}), handlers.Login)
	m.Get("/logout", binding.Json(response.Attribute{}), handlers.Logout)

	m.Run()
}
