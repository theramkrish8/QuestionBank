package models

import "gopkg.in/mgo.v2/bson"

type (
	//Class object
	Class struct {
		ClassID  bson.ObjectId `json:class_id bson:"class_id,omitempty"`
		Standard string        `json:"standard" bson:"standard"`
		Section  string        `json:"sections" bson:"sections"`
		Subject  Subject       `json:"subjects" bson:"subjects"`
	}
)
