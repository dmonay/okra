package handlers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (dw *DoWorkResource) UpdateObjProperties(c *gin.Context) {
	org := c.Params.ByName("organization")
	treeId := c.Params.ByName("treeid")
	obj := c.Params.ByName("objective")
	id := bson.ObjectIdHex(treeId)

	type ObjPropertiesJson struct {
		ObjName   string      `json:"objName",omitempty`
		ObjBody   string      `json:"objbody",omitempty`
		Completed interface{} `json:"completed",omitempty`
	}

	var reqBody ObjPropertiesJson
	c.Bind(&reqBody)
	newName := reqBody.ObjName
	newBody := reqBody.ObjBody
	newStatus := reqBody.Completed

	// instead of using bson.M, manually create the map
	// so that we can add fields conditionally
	myMap := make(map[string]interface{})

	nameKey := "objectives.$.name"
	bodyKey := "objectives.$.body"
	statusKey := "objectives.$.completed"

	if newName != "" {
		myMap[nameKey] = newName
	}

	if newBody != "" {
		myMap[bodyKey] = newBody
	}

	if newStatus != "" {
		myMap[statusKey] = newStatus
	}

	querier := bson.M{"_id": id, "objectives.id": obj}
	updateName := bson.M{"$set": myMap}
	err := dw.mongo.C(org).Update(querier, updateName)
	CheckErr(err, "Mongo failed to update objective's properties")

	c.JSON(201, "Success")
}
