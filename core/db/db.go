package db

import (
	"sync"

	"gopkg.in/mgo.v2"
)

//DBController type to have mongodb session object
type DBController struct {
	session *mgo.Session
}

var instance *DBController
var once sync.Once

//GetInstance is to get singleton instance of db controller
func GetInstance() *DBController {
	once.Do(func() {
		s, err := mgo.Dial("mongodb://localhost")
		// Check if connection error, is mongo running?
		if err != nil {
			panic(err)
		}
		instance = &DBController{}
		instance.session = s
	})
	return instance
}

//GetSession to get the db session to add data to db
func (dc DBController) GetSession() *mgo.Session {
	if dc.session == nil {
		panic("Uninitialized session!!!")
	}
	return dc.session
}

const (
	//DBName Mongo DB Name
	DBName = "QuestionBank"
	//UserCollection User collection Name
	UserCollection = "users"
	//GroupCollection group collection name
	GroupCollection = "groups"
	//SchoolCollection schools collection name
	SchoolCollection = "schools"
	//ClassCollection class collection name
	ClassCollection = "classes"
)
