package handlers

import (
	"fmt"
	// "github.com/dmonay/do-work-api/authentication"
	"github.com/dmonay/do-work-api/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/cookiejar"
)

func (dw *DoWorkResource) Create(c *gin.Context) {

	var creds common.Credentials

	c.Bind(&creds)

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}
	fmt.Println("cookie:", client)

	c.JSON(200, "You have successfully created an item")

}
