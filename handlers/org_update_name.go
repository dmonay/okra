package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) UpdateOrgName(c *gin.Context) {

	var reqBody common.OrgUpdateNameJson
	c.Bind(&reqBody)

	oldName := reqBody.OldName
	newName := reqBody.NewName

	var result interface{}
	err := dw.mongo.Run(bson.D{{"renameCollection", oldName}, {"to", newName}}, &result)

	if err != nil {
		CheckErr(err, "Mongo failed to update org name", c)
		return
	}

	c.JSON(201, SuccessMsg{result})

	// The admin database is a special database inside of mongo.
	// It is where mongo will store administrative settings like users and what they have access to.
	// It isn't created by default (afaik), it's just required that you can access it to run admin commands.
}
