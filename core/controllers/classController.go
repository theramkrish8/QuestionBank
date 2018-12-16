package controllers

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/db"
	"QuestionBank/core/models"
)

//ClassController type handles class collection operations
type ClassController struct {
}

var classInstance *ClassController
var classOnce sync.Once

//GetClassController returns singleton instance of ClassController class
func GetClassController() *ClassController {
	classOnce.Do(func() {
		classInstance = &ClassController{}
	})
	return classInstance
}

//GetAllClasses returns all classes list from ClassCollection
func (uc ClassController) GetAllClasses() ([]models.Class, error) {
	var results []models.Class
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.ClassCollection).Find(nil).All(&results)
	return results, err
}

//GetClassByID returns particular Class by id
func (uc ClassController) GetClassByID(oid bson.ObjectId) (models.Class, error) {
	u := models.Class{}
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.ClassCollection).Find(oid).One(&u)
	return u, err
}

//InsertClass insert new Class to Class collection
func (uc ClassController) InsertClass(class models.Class) (models.Class, error) {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.ClassCollection).Insert(&class)
	return class, err
}

//DeleteClass deletes class from Classcollection
func (uc ClassController) DeleteClass(oid bson.ObjectId) error {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.ClassCollection).RemoveId(oid)
	return err
}

//UpdateClassByID updates class data identified by Class
func (uc ClassController) UpdateClassByID(oid bson.ObjectId, class models.Class) (models.Class, error) {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.ClassCollection).UpdateId(&oid, &class)
	return class, err
}
