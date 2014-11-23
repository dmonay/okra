package handlers

import (
	// "fmt"
	"github.com/codegangsta/martini-contrib/binding"
	// "github.com/dmonay/do-work-api/common"
	// "github.com/dmonay/do-work-api/database"
	// "github.com/coopernurse/gorp"
	// "fmt"
	"net/http"
)

func Register(attr Credentials, erro binding.Errors) (int, string) {
	uname := attr.Username

	// fmt.Println("db:", database.Dbmap)

	// query := &common.User{0, uname, attr.Password}
	// err = dbmap.Insert(query)

	// pwd := attr.Password
	return http.StatusOK, JsonString(SuccessMsg{"You have successfully registered, " + uname})
}
