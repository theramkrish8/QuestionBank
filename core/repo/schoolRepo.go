package repo

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/mgorepo"
	"QuestionBank/core/models"
)

//SchoolRepo type handles user collection operations
type SchoolRepo struct {
}

var schoolInstance *SchoolRepo
var schoolOnce sync.Once

//GetSchoolRepo returns singleton instance of SchoolRepo class
func GetSchoolRepo() *SchoolRepo {
	schoolOnce.Do(func() {
		schoolInstance = &SchoolRepo{}
	})
	return schoolInstance
}

//GetAllSchools returns all users list from SchoolCollection
func (sc SchoolRepo) GetAllSchools() ([]models.School, error) {
	var results []models.School
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.SchoolCollection).Find(nil).All(&results)
	return results, err
}

//GetSchoolByID returns particular school by id
func (sc SchoolRepo) GetSchoolByID(oid bson.ObjectId) (models.School, error) {
	s := models.School{}
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.SchoolCollection).Find(oid).One(&s)
	return s, err
}

//InsertSchool insert new school to school collection
func (sc SchoolRepo) InsertSchool(s models.School) (models.School, error) {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.SchoolCollection).Insert(&s)
	return s, err
}

//DeleteSchool deletes school from schoolcollection
func (sc SchoolRepo) DeleteSchool(oid bson.ObjectId) error {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.SchoolCollection).RemoveId(oid)
	return err
}

//UpdateSchoolByID updates school identified by school id
func (sc SchoolRepo) UpdateSchoolByID(s models.School) (models.School, error) {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.SchoolCollection).UpdateId(s.SchoolID, &s)
	return s, err
}
