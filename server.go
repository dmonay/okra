package main

import (
	// "fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/dmonay/do-work-api/handlers"
	"github.com/dmonay/do-work-api/middleware"
)

func main() {

	//initialize mysql
	dbmap := middleware.InitDb()
	defer dbmap.Db.Close()

	m := martini.Classic()

	//define the endpoints
	m.Post("/register", binding.Json(handlers.Credentials{}), handlers.Register)
	m.Post("/login", binding.Json(handlers.Credentials{}), handlers.Login)
	m.Post("/logout", binding.Json(handlers.Credentials{}), handlers.Logout)

	m.Run()
}
