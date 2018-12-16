package controllers

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/db"
	"QuestionBank/core/models"
)

//SchoolController type handles user collection operations
type SchoolController struct {
}

var schoolInstance *SchoolController
var schoolOnce sync.Once

//GetSchoolController returns singleton instance of SchoolController class
func GetSchoolController() *SchoolController {
	schoolOnce.Do(func() {
		schoolInstance = &SchoolController{}
	})
	return schoolInstance
}

//GetAllSchools returns all users list from SchoolCollection
func (sc SchoolController) GetAllSchools() ([]models.School, error) {
	var results []models.School
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.SchoolCollection).Find(nil).All(&results)
	return results, err
}

//GetSchoolByID returns particular school by id
func (sc SchoolController) GetSchoolByID(oid bson.ObjectId) (models.School, error) {
	s := models.School{}
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.SchoolCollection).Find(oid).One(&s)
	return s, err
}

//InsertSchool insert new school to school collection
func (sc SchoolController) InsertSchool(s models.School) (models.School, error) {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.SchoolCollection).Insert(&s)
	return s, err
}

//DeleteSchool deletes school from schoolcollection
func (sc SchoolController) DeleteSchool(oid bson.ObjectId) error {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.SchoolCollection).RemoveId(oid)
	return err
}

//UpdateSchoolByID updates school identified by school id
func (sc SchoolController) UpdateSchoolByID(oid bson.ObjectId, s models.School) (models.School, error) {
	err := db.GetInstance().GetSession().DB(db.DBName).C(db.SchoolCollection).UpdateId(&oid, &s)
	return s, err
}
