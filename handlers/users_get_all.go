package handlers

import (
	"github.com/dmonay/okra/authentication"
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func (dw *DoWorkResource) GetAllUsers(c *gin.Context) {
	user := c.Params.ByName("user")

	var results []common.UsersObj
	err := dw.mongo.C("Users").Find(bson.M{"username": bson.M{"$regex": bson.RegEx{Pattern: user, Options: "i"}}}).All(&results)
	if err != nil {
		CheckErr(err, "Failed to get all users", c)
		return
	}
	result := make(map[string][]byte)
	count := 0
	for _, value := range results {
		count++
		rawKey := os.Getenv("SECRET")
		key := []byte(rawKey) // 32 bytes
		plaintext := []byte(value.Id)

		ciphertext, err2 := authentication.Encrypt(key, plaintext)
		if err2 != nil {
			CheckErr(err2, "Failed to encrypt key", c)
			return
		}
		result[value.Username] = ciphertext
		if count == 7 {
			break
		}
	}

	c.JSON(200, SuccessMsg{result})
}
