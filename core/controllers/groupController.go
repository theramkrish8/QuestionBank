package controllers

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/db"
	"QuestionBank/core/models"
)

//GroupController class
type GroupController struct {
}

var groupInstance *GroupController
var groupOnce sync.Once

//GetGroupController - return singleton instance of GroupController
func GetGroupController() *GroupController {
	groupOnce.Do(func() {
		groupInstance = &GroupController{}
	})
	return groupInstance
}

//GetAllGroups returns all the groups
func (uc GroupController) GetAllGroups() ([]models.Group, error) {
	var results []models.Group
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.GroupCollection).Find(nil).All(&results)
	return results, err
}

//GetGroupByID return particular group identified by group id.
func (uc GroupController) GetGroupByID(oid bson.ObjectId) (models.Group, error) {
	u := models.Group{}
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.GroupCollection).Find(oid).One(&u)
	return u, err
}

//InsertGroup inserts new group
func (uc GroupController) InsertGroup(group models.Group) (models.Group, error) {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.GroupCollection).Insert(&group)
	return group, err
}

//DeleteGroup deletes group
func (uc GroupController) DeleteGroup(oid bson.ObjectId) error {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.GroupCollection).RemoveId(oid)
	return err
}

//UpdateGroupWithUserByID adds user to group
func (uc GroupController) UpdateGroupWithUserByID(oid bson.ObjectId, uid bson.ObjectId) error {
	change := bson.M{"$push": bson.M{"users": uid}}
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.GroupCollection).UpdateId(oid, change)
	return err
}

//UpdateGroupWithSchoolByID adds school to group
func (uc GroupController) UpdateGroupWithSchoolByID(gid bson.ObjectId, sid bson.ObjectId) error {
	//schoolRef := bson.DBPointer{Namespace: "schools", Id: sid}
	change := bson.M{"$push": bson.M{"schools": sid}}
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.GroupCollection).UpdateId(gid, change)
	return err
}
