package services

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/models"
	"QuestionBank/core/repo"
	"QuestionBank/core/utils"
)

// SchoolList Get all users
func SchoolList(c *gin.Context) {
	var results, err = repo.GetSchoolRepo().GetAllSchools()
	if err != nil {
		utils.CheckAndDisplay(err, "Schools doesn't exist", "404", c)
	}
	c.JSON(200, results)
}

// GetSchool retrieves an individual user resource
func GetSchool(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the id to ObjectId
	oid := bson.ObjectIdHex(id)
	//Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("School ID is not a bson.ObjectId", "404", c)
		return
	}

	var s, err = repo.GetSchoolRepo().GetSchoolByID(oid)
	// Fetch user
	if err != nil {
		utils.DisplayError("Users doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, s)
}

// CreateSchool creates a new user resource
func CreateSchool(c *gin.Context) {

	var json models.School

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	var s, err = repo.GetSchoolRepo().InsertSchool(json)

	if err != nil {
		utils.CheckAndDisplay(err, "Insert failed", "403", c)
	} else {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, s)
	}
}

// RemoveSchool removes an existing user resource
func RemoveSchool(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Convert the user id to objectId
	oid := bson.ObjectIdHex(id)

	// Verify id is ObjectId, otherwise bail
	if !oid.Valid() {
		utils.DisplayError("School ID is not a bson.ObjectId", "404", c)
		return
	}

	// Remove school
	var err = repo.GetSchoolRepo().DeleteSchool(oid)
	if err != nil {
		utils.CheckAndDisplay(err, "Fail to Remove", "404", c)
		return
	}
	utils.MessageTypeDefault("Success", c)
}

// UpdateSchool Updates an existing user resource
func UpdateSchool(c *gin.Context) {
	var school models.School

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&school)
	// Verify id is ObjectId, otherwise bail
	if !school.SchoolID.Valid() {
		utils.DisplayError("ID is not a bson.ObjectId", "404", c)
		return
	}

	// Grab id

	var s, err = repo.GetSchoolRepo().UpdateSchoolByID(school)
	if err != nil {
		utils.CheckAndDisplay(err, "Update failed", "403", c)
	} else {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, s)
	}

}
