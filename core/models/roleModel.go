package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/enums"
)

type (
	//Role type
	Role struct {
		RoleID     bson.ObjectId  `json:"role_id"  bson:"_id,omitempty"`
		Role       enums.RoleType `json:"role" bson:"role"`
		Class      mgo.DBRef      `json:"classes" bson:"classes"`
		AssignedOn string         `json:"assigned_on" bson:"assigned_on"`
	}
)
