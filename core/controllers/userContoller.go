package controllers

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/db"
	"QuestionBank/core/models"
)

//UserController type handles user collection operations
type UserController struct {
}

var userInstance *UserController
var userOnce sync.Once

//GetUserController returns singleton instance of UserController class
func GetUserController() *UserController {
	userOnce.Do(func() {
		userInstance = &UserController{}
	})
	return userInstance
}

//GetAllUsers returns all users list from UserCollection
func (uc UserController) GetAllUsers() ([]models.User, error) {
	var results []models.User
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.UserCollection).Find(nil).All(&results)
	return results, err
}

//GetUserByID returns particular User by id
func (uc UserController) GetUserByID(oid bson.ObjectId) (models.User, error) {
	u := models.User{}
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.UserCollection).Find(oid).One(&u)
	return u, err
}

//InsertUser insert new user to user collection
func (uc UserController) InsertUser(user models.User) (models.User, error) {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.UserCollection).Insert(&user)
	return user, err
}

//DeleteUser deletes user from usercollection
func (uc UserController) DeleteUser(oid bson.ObjectId) error {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.UserCollection).RemoveId(oid)
	return err
}

//UpdateUserByID updates user data identified by user
func (uc UserController) UpdateUserByID(oid bson.ObjectId, user models.User) (models.User, error) {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.UserCollection).UpdateId(&oid, &user)
	return user, err
}
