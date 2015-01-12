package handlers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) UpdateKrProperties(c *gin.Context) {
	org := c.Params.ByName("organization")
	treeId := c.Params.ByName("treeid")
	obj := c.Params.ByName("objective")
	kr := c.Params.ByName("kr")
	id := bson.ObjectIdHex(treeId)

	type KrPropertiesJson struct {
		KrName    string      `json:"krName",omitempty`
		KrBody    string      `json:"krbody",omitempty`
		Priority  string      `json:"priority",omitempty`
		Completed interface{} `json:"completed",omitempty`
	}

	var reqBody KrPropertiesJson
	c.Bind(&reqBody)
	newName := reqBody.KrName
	newBody := reqBody.KrBody
	newStatus := reqBody.Completed
	newPriority := reqBody.Priority

	// instead of using bson.M, manually create the map
	// so that we can add fields conditionally
	myMap := make(map[string]interface{})

	// TEMP workaround. Mongo doesn't support >1 positional operators per query,
	// so I can't use the index of the nested array inside the objectives document.
	// Thus I need the index of the array beforehand, and this is currently,
	// temporarily passed in as the last URL param
	nameKey := "objectives.$.keyresults." + kr + ".name"
	bodyKey := "objectives.$.keyresults." + kr + ".body"
	statusKey := "objectives.$.keyresults." + kr + ".completed"
	prioritiesKey := "objectives.$.keyresults." + kr + ".priority"

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
		CheckErr(err, "Mongo failed to update key result's properties", c)
		return
	}

	c.JSON(201, SuccessMsg{"Key result updated!"})
}
