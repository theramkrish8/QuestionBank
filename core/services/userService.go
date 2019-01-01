package services

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/models"
	"QuestionBank/core/repo"
	"QuestionBank/core/utils"
)

// UsersList Get all users
func UsersList(c *gin.Context) {
	var results, err = repo.GetUserRepo().GetAllUsers()
	if err != nil {
		utils.CheckAndDisplay(err, "Users doesn't exist", "404", c)
	}
	c.JSON(200, results)
}

// GetUser retrieves an individual user resource
func GetUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the id to ObjectId
	oid := bson.ObjectIdHex(id)
	//Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("ID is not a bson.ObjectId", "404", c)
		return
	}

	var u, err = repo.GetUserRepo().GetUserByID(oid)
	// Fetch user
	if err != nil {
		utils.DisplayError("Users doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, u)
}

// CreateUser creates a new user resource
func CreateUser(c *gin.Context) {

	var json models.User

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	var user, err = repo.GetUserRepo().InsertUser(json)

	if err != nil {
		utils.CheckAndDisplay(err, "Insert user failed", "403", c)
	} else {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, user)
	}
}

// RemoveUser removes an existing user resource
func RemoveUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the user id to objectId
	oid := bson.ObjectIdHex(id)

	// Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("ID is not a bson.ObjectId", "404", c)
		return
	}

	// Remove user
	var err = repo.GetUserRepo().DeleteUser(oid)
	if err != nil {
		utils.CheckAndDisplay(err, "Fail to Remove user", "404", c)
		return
	}
	utils.MessageTypeDefault("Success", c)
}

// UpdateUser Updates an existing user resource
func UpdateUser(c *gin.Context) {
	var user models.User

	c.Bind(&user)

	// Verify id is ObjectId, otherwise fail
	if !user.UserID.Valid() {
		utils.DisplayError("User ID is not a bson.ObjectId", "404", c)
		return
	}

	var u, err = repo.GetUserRepo().UpdateUserByID(user)
	if err != nil {
		utils.CheckAndDisplay(err, "Update failed", "403", c)
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, u)
}
