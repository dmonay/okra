package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) GetAllTrees(c *gin.Context) {
	org := c.Params.ByName("organization")

	var intermResult []common.OkrTree
	err := dw.mongo.C(org).Find(bson.M{"type": "tree"}).All(&intermResult)
	if err != nil {
		CheckErr(err, "Failed to retrieve trees in organization "+org+" from Mongo", c)
		return
	}

	length := len(intermResult)
	result := make([]common.TreeInOrg, length)

	for key, value := range intermResult {
		treeName := value.TreeName
		treeId := value.Id
		active := value.Active

		result[key].Name = treeName
		result[key].Id = treeId
		result[key].Active = active
	}

	c.JSON(200, SuccessMsg{result})
}
