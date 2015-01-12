package handlers

import (
	"github.com/dmonay/okra/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) UpdateTaskProperties(c *gin.Context) {
	org := c.Params.ByName("organization")
	treeId := c.Params.ByName("treeid")
	obj := c.Params.ByName("objective")
	krIndex := c.Params.ByName("kr")
	taskIndex := c.Params.ByName("task")
	id := bson.ObjectIdHex(treeId)

	var reqBody common.TaskPropertiesJson
	c.Bind(&reqBody)
	newName := reqBody.TaskName
	newBody := reqBody.TaskBody
	newStatus := reqBody.Completed
	newPriority := reqBody.Priority

	// instead of using bson.M, manually create the map
	// so that we can add fields conditionally
	myMap := make(map[string]interface{})

	// TEMP workaround. Mongo doesn't support >1 positional operators per query,
	// so I can't use the index of the nested array inside the objectives document.
	// Thus I need the index of the array beforehand, and this is currently,
	// temporarily passed in as the last URL param
	nameKey := "objectives.$.keyresults." + krIndex + ".tasks." + taskIndex + ".name"
	bodyKey := "objectives.$.keyresults." + krIndex + ".tasks." + taskIndex + ".body"
	statusKey := "objectives.$.keyresults." + krIndex + ".tasks." + taskIndex + ".completed"
	prioritiesKey := "objectives.$.keyresults." + krIndex + ".tasks." + taskIndex + ".priority"

	if newName != "" {
		myMap[nameKey] = newName
	}

	if newBody != "" {
		myMap[bodyKey] = newBody
	}

	if newStatus != "" {
		myMap[statusKey] = newStatus
	}

	if newPriority != "" {
		myMap[prioritiesKey] = newPriority
	}

	querier := bson.M{"_id": id, "objectives.id": obj}
	updateName := bson.M{"$set": myMap}
	err := dw.mongo.C(org).Update(querier, updateName)
	if err != nil {
		CheckErr(err, "Mongo failed to update task's properties", c)
		return
	}

	c.JSON(201, SuccessMsg{"Task updated!"})
}
