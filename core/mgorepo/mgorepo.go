package mgorepo

import (
	"sync"

	"gopkg.in/mgo.v2"
)

//DBRepo type to have mongodb session object
type DBRepo struct {
	session *mgo.Session
}

var instance *DBRepo
var once sync.Once

//GetInstance is to get singleton instance of db Repo
func GetInstance() *DBRepo {
	once.Do(func() {
		s, err := mgo.Dial("mongodb://localhost")
		// Check if connection error, is mongo running?
		if err != nil {
			panic(err)
		}
		instance = &DBRepo{}
		instance.session = s
	})
	return instance
}

//GetSession to get the db session to add data to db
func (dc DBRepo) GetSession() *mgo.Session {
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
