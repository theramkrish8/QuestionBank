package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type TeacherDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "questionbank"
)

func (m *TeacherDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
