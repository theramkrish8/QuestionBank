package repo

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/mgorepo"
	"QuestionBank/core/models"
)

//GroupRepo class
type GroupRepo struct {
}

var groupInstance *GroupRepo
var groupOnce sync.Once

//GetGroupRepo - return singleton instance of GroupRepo
func GetGroupRepo() *GroupRepo {
	groupOnce.Do(func() {
		groupInstance = &GroupRepo{}
	})
	return groupInstance
}

//GetAllGroups returns all the groups
func (uc GroupRepo) GetAllGroups() ([]models.Group, error) {
	var results []models.Group
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.GroupCollection).Find(nil).All(&results)
	return results, err
}

//GetGroupByID return particular group identified by group id.
func (uc GroupRepo) GetGroupByID(oid bson.ObjectId) (models.Group, error) {
	u := models.Group{}
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.GroupCollection).Find(oid).One(&u)
	return u, err
}

//InsertGroup inserts new group
func (uc GroupRepo) InsertGroup(group models.Group) (models.Group, error) {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.GroupCollection).Insert(&group)
	return group, err
}

//DeleteGroup deletes group
func (uc GroupRepo) DeleteGroup(oid bson.ObjectId) error {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.GroupCollection).RemoveId(oid)
	return err
}

//UpdateGroupWithUserByID adds user to group
func (uc GroupRepo) UpdateGroupWithUserByID(oid bson.ObjectId, uid bson.ObjectId) error {
	change := bson.M{"$push": bson.M{"users": uid}}
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.GroupCollection).UpdateId(oid, change)
	return err
}

//UpdateGroupWithSchoolByID adds school to group
func (uc GroupRepo) UpdateGroupWithSchoolByID(gid bson.ObjectId, sid bson.ObjectId) error {
	//schoolRef := bson.mgorepoPointer{Namespace: "schools", Id: sid}
	change := bson.M{"$push": bson.M{"schools": sid}}
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.GroupCollection).UpdateId(gid, change)
	return err
}
