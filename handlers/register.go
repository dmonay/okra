package handlers

import (
	// "fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"net/http"
)

func Register(attr Credentials, err binding.Errors) (int, string) {
	uname := attr.Username

	// pwd := attr.Password
	return http.StatusOK, JsonString(SuccessMsg{"You have successfully registered, " + uname})
}
