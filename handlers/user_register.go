package handlers

import (
	"github.com/dmonay/okra/authentication"
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func (dw *DoWorkResource) Register(c *gin.Context) {

	var user common.UserJson
	c.Bind(&user)

	uname := user.Username
	first := user.First
	last := user.Last
	ga := user.GaUser
	orgs := []string{}
	trees := []string{}
	id := bson.NewObjectId()

	err := dw.mongo.C("Users").Insert(&common.UsersObj{uname, orgs, trees, first, last, ga, id})
	if err != nil {
		CheckErr(err, "User not added to Users collection", c)
		return
	}

	rawKey := os.Getenv("SECRET")
	key := []byte(rawKey) // 32 bytes
	plaintext := []byte(id)

	ciphertext, err2 := authentication.Encrypt(key, plaintext)
	if err2 != nil {
		CheckErr(err2, "Failed to encrypt key", c)
		return
	}

	c.JSON(201, SuccessMsg{ciphertext})
}
