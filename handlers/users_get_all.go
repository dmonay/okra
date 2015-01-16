package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) GetAllUsers(c *gin.Context) {
	user := c.Params.ByName("user")

	var results []common.UsersObj
	err := dw.mongo.C("Users").Find(bson.M{"username": bson.M{"$regex": bson.RegEx{Pattern: user, Options: "i"}}}).All(&results)
	if err != nil {
		CheckErr(err, "Failed to get all users", c)
		return
	}
	var result []string
	count := 0
	for _, value := range results {
		count++
		result = append(result, value.Username)
		if count == 7 {
			break
		}
	}

	c.JSON(200, SuccessMsg{result})
}
