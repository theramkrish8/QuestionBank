package services

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/controllers"
	"QuestionBank/core/models"
	"QuestionBank/core/utils"
)

// GroupList Get all users
func GroupList(c *gin.Context) {
	var results, err = controllers.GetGroupController().GetAllGroups()
	if err != nil {
		utils.CheckAndDisplay(err, "Group doesn't exist", "404", c)
	}
	c.JSON(200, results)
}

// GetGroup retrieves an individual user resource
func GetGroup(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the id to ObjectId
	oid := bson.ObjectIdHex(id)
	//Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("ID is not a bson.ObjectId", "404", c)
		return
	}

	var u, err = controllers.GetGroupController().GetGroupByID(oid)
	// Fetch group
	if err != nil {
		utils.DisplayError("Group doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, u)
}

// CreateGroup creates a new user resource
func CreateGroup(c *gin.Context) {

	var json models.Group

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	var group, err = controllers.GetGroupController().InsertGroup(json)

	if err != nil {
		utils.CheckAndDisplay(err, "Insert group failed", "403", c)
	} else {
		if group.GroupName == json.GroupName {
			content := gin.H{
				"result":    "Success",
				"GroupName": group.GroupName,
				"GroupCode": group.GroupCode,
				"Board":     group.Board,
				"Users":     group.Users,
				"Schools":   group.Schools,
				"Classes":   group.Classes,
			}
			c.Writer.Header().Set("Content-Type", "application/json")
			c.JSON(201, content)
		}
	}
}

// RemoveGroup removes an existing user resource
func RemoveGroup(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the user id to objectId
	oid := bson.ObjectIdHex(id)

	// Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("Group ID is not a bson.ObjectId", "404", c)
		return
	}

	// Remove user
	var err = controllers.GetGroupController().DeleteGroup(oid)
	if err != nil {
		utils.CheckAndDisplay(err, "Fail to Remove", "404", c)
		return
	}
	utils.MessageTypeDefault("Success", c)
}

// AddUserToGroup Updates an existing user resource
func AddUserToGroup(c *gin.Context) {
	// Grab id
	gid := c.Params.ByName("id1")
	uid := c.Params.ByName("id1")
	//Convert id to bson.ObjectId
	goid := bson.ObjectIdHex(gid)
	uoid := bson.ObjectIdHex(uid)

	// Verify id is ObjectId, otherwise bail
	if !goid.Valid() && !uoid.Valid() {
		utils.DisplayError("ID is not a bson.ObjectId", "404", c)
		return
	}

	// Grab id

	err := controllers.GetGroupController().UpdateGroupWithUserByID(goid, uoid)
	if err != nil {
		utils.CheckAndDisplay(err, "Update failed", "403", c)
	}
	content := gin.H{
		"result": "Success",
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)

}

// AddSchoolToGroup Updates an existing user resource
func AddSchoolToGroup(c *gin.Context) {
	// Grab id
	goid := c.Params.ByName("id1")
	soid := c.Params.ByName("id2")
	//Convert id to bson.ObjectId
	gid := bson.ObjectIdHex(goid)
	sid := bson.ObjectIdHex(soid)

	// Verify id is ObjectId, otherwise bail
	if !gid.Valid() && !sid.Valid() {
		utils.DisplayError("School ID or  is not a bson.ObjectId", "404", c)
		return
	}

	// Grab id

	var err = controllers.GetGroupController().UpdateGroupWithSchoolByID(gid, sid)
	if err != nil {
		utils.CheckAndDisplay(err, "Update failed", "403", c)
	}
	content := gin.H{
		"result": "Success",
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)

}
