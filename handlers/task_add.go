package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) CreateTask(c *gin.Context) {

	org := c.Params.ByName("organization")
	obj := c.Params.ByName("objective")
	krIndex := c.Params.ByName("kr")
	var reqBody common.TaskJson

	c.Bind(&reqBody)

	arrayOfMembers := []common.Member{}
	for _, value := range reqBody.Members {
		// Get userId of user
		var result common.UsersObj
		err := dw.mongo.C("Users").Find(bson.M{"username": value.Username}).One(&result)
		if err != nil {
			CheckErr(err, "Mongo failed to find the "+value.Username+"'s doc in Users", c)
			return
		}
		memObj := common.Member{value.Username, result.Id.Hex(), value.Role}
		arrayOfMembers = append(arrayOfMembers, memObj)
	}

	treeId := bson.ObjectIdHex(reqBody.TreeId)

	kr := common.TasksModel{
		Id:        reqBody.Id,
		Name:      reqBody.Name,
		Body:      reqBody.Body,
		Completed: reqBody.Completed,
		Members:   arrayOfMembers,
		Priority:  reqBody.Priority,
	}

	// again, because Mongo doesn't allow for >1 positional operator,
	// we must receive the index of the key result from the client and
	// use that in our query
	findKR := "objectives.$.keyresults." + krIndex + ".tasks"

	colQuerier := bson.M{"_id": treeId, "objectives.id": obj}
	addKeyResult := bson.M{"$push": bson.M{findKR: kr}}
	err := dw.mongo.C(org).Update(colQuerier, addKeyResult)
	if err != nil {
		CheckErr(err, "Mongo failed to add task", c)
		return
	}

	c.JSON(201, SuccessMsg{"You have successfully added a task to the key result"})
}
