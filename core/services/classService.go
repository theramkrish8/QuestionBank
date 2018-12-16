package services

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/controllers"
	"QuestionBank/core/models"
	"QuestionBank/core/utils"
)

// ClassList Get all classes
func ClassList(c *gin.Context) {
	var results, err = controllers.GetClassController().GetAllClasses()
	if err != nil {
		utils.CheckAndDisplay(err, "Classes doesn't exist", "404", c)
	}
	c.JSON(200, results)
}

// GetClass retrieves an individual class resource
func GetClass(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the id to ObjectId
	oid := bson.ObjectIdHex(id)
	//Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("ID is not a bson.ObjectId", "404", c)
		return
	}

	var u, err = controllers.GetClassController().GetClassByID(oid)
	// Fetch Class
	if err != nil {
		utils.DisplayError("Classes doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, u)
}

// CreateClass creates a new Class resource
func CreateClass(c *gin.Context) {

	var json models.Class

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	var class, err = controllers.GetClassController().InsertClass(json)

	if err != nil {
		utils.CheckAndDisplay(err, "Insert Class failed", "403", c)
	} else {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, class)
	}
}

// RemoveClass removes an existing class resource
func RemoveClass(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the class id to objectId
	oid := bson.ObjectIdHex(id)

	// Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("ID is not a bson.ObjectId", "404", c)
		return
	}

	// Remove Class
	var err = controllers.GetClassController().DeleteClass(oid)
	if err != nil {
		utils.CheckAndDisplay(err, "Fail to Remove Class", "404", c)
		return
	}
	utils.MessageTypeDefault("Success", c)
}

// UpdateClass Updates an existing Class resource
func UpdateClass(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	//Convert id to bson.ObjectId
	oid := bson.ObjectIdHex(id)
	var json models.Class

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	// Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("Class ID is not a bson.ObjectId", "404", c)
		return
	}

	// Grab id

	var u, err = controllers.GetClassController().UpdateClassByID(oid, json)
	if err != nil {
		utils.CheckAndDisplay(err, "Update failed", "403", c)
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, u)
}
