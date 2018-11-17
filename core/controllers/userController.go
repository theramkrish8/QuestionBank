package controllers

import (
	"log"

	"../models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

const (
	//DBName Mongo DB Name
	DBName = "QuestionBank"
	//DBCollection Collection Name
	DBCollection = "users"
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func messageTypeDefault(msg string, c *gin.Context) {
	content := gin.H{
		"status": "200",
		"result": msg,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func checkErrTypeOne(err error, msg string, status string, c *gin.Context) {
	if err != nil {
		log.Fatalln(msg, err)
		content := gin.H{
			"status": status,
			"result": msg,
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(200, content)
		panic(err)
	}
}

func checkErrTypeTwo(msg string, status string, c *gin.Context) {
	content := gin.H{
		"status": status,
		"result": msg,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}

// UsersList Get all users
func (uc UserController) UsersList(c *gin.Context) {

	var results []models.User
	err := uc.session.DB(DBName).C(DBCollection).Find(nil).All(&results)
	if err != nil {
		checkErrTypeOne(err, "Users doesn't exist", "404", c)
		return
	}

	c.JSON(200, results)
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	//Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.User{}
	err := uc.session.DB(DBName).C(DBCollection).FindId(oid).One(&u)
	// Fetch user
	if err != nil {
		checkErrTypeTwo("Users doesn't exist", "404", c)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, u)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(c *gin.Context) {

	var json models.User

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	u := uc.createUser(json.FirstName, json.LastName, json.Gender, json.Mobile, json.Email, json.Password, c)
	if u.FirstName == json.FirstName {
		content := gin.H{
			"result":    "Success",
			"FirstName": u.FirstName,
			"LastName":  u.LastName,
			"Gender":    u.Gender,
			"Mobile":    u.Mobile,
			"Email":     u.Email,
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.session.DB(DBName).C(DBCollection).RemoveId(oid); err != nil {
		checkErrTypeOne(err, "Fail to Remove", "404", c)
		return
	}

	messageTypeDefault("Success", c)

}

// UpdateUser Updates an existing user resource
func (uc UserController) UpdateUser(c *gin.Context) {
	// Grab id
	id := c.Params.ByName("id")
	var json models.User

	// This will infer what binder to use depending on the content-type header.
	c.Bind(&json)

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		checkErrTypeTwo("ID is not a bson.ObjectId", "404", c)
		return
	}

	// Grab id

	u := uc.updateUser(id, json.FirstName, json.LastName, json.Gender, json.Mobile, json.Email, json.Password, c)

	if u.FirstName == json.FirstName {
		content := gin.H{
			"result":    "Success",
			"FirstName": u.FirstName,
			"LastName":  u.LastName,
			"Gender":    u.Gender,
			"Mobile":    u.Mobile,
			"Email":     u.Email,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}

}

func (uc UserController) createUser(FirstName string, LastName string, Gender string, Mobile string, Email string, Password string, c *gin.Context) models.User {
	user := models.User{
		FirstName: FirstName,
		LastName:  LastName,
		Gender:    Gender,
		Mobile:    Mobile,
		Email:     Email,
		Password:  Password,
	}
	// Write the user to mongo
	err := uc.session.DB(DBName).C(DBCollection).Insert(&user)
	checkErrTypeOne(err, "Insert failed", "403", c)
	return user
}

func (uc UserController) updateUser(ID string, FirstName string, LastName string, Gender string, Mobile string, Email string, Password string, c *gin.Context) models.User {

	user := models.User{
		FirstName: FirstName,
		LastName:  LastName,
		Gender:    Gender,
		Mobile:    Mobile,
		Email:     Email,
		Password:  Password,
	}

	oid := bson.ObjectIdHex(ID)
	// Write the user to mongo
	if err := uc.session.DB(DBName).C(DBCollection).UpdateId(oid, &user); err != nil {
		checkErrTypeOne(err, "Update failed", "403", c)

	}

	return user
}
