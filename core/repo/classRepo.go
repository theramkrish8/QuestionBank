package repo

import (
	"sync"

	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/mgorepo"
	"QuestionBank/core/models"
)

//ClassRepo type handles class collection operations
type ClassRepo struct {
}

var classInstance *ClassRepo
var classOnce sync.Once

//GetClassRepo returns singleton instance of ClassRepo class
func GetClassRepo() *ClassRepo {
	classOnce.Do(func() {
		classInstance = &ClassRepo{}
	})
	return classInstance
}

//GetAllClasses returns all classes list from ClassCollection
func (uc ClassRepo) GetAllClasses() ([]models.Class, error) {
	var results []models.Class
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.ClassCollection).Find(nil).All(&results)
	return results, err
}

//GetClassByID returns particular Class by id
func (uc ClassRepo) GetClassByID(oid bson.ObjectId) (models.Class, error) {
	u := models.Class{}
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.ClassCollection).Find(oid).One(&u)
	return u, err
}

//InsertClass insert new Class to Class collection
func (uc ClassRepo) InsertClass(class models.Class) (models.Class, error) {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.ClassCollection).Insert(&class)
	return class, err
}

//DeleteClass deletes class from Classcollection
func (uc ClassRepo) DeleteClass(oid bson.ObjectId) error {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.ClassCollection).RemoveId(oid)
	return err
}

//UpdateClassByID updates class data identified by Class
func (uc ClassRepo) UpdateClassByID(oid bson.ObjectId, class models.Class) (models.Class, error) {
	err := mgorepo.GetInstance().GetSession().DB(mgorepo.DBName).C(mgorepo.ClassCollection).UpdateId(&oid, &class)
	return class, err
}
