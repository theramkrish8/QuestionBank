package models

import (
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/enums"
)

type (
	//User object
	User struct {
		UserID    bson.ObjectId `json:"user_id" bson:"_id,omitempty"`
		FirstName string        `json:"first_name" bson:"first_name"`
		LastName  string        `json:"last_name" bson:"last_name"`
		Gender    string        `json:"gender" bson:"gender"`
		Mobile    string        `json:"mobile" bson:"mobile"`
		Email     string        `json:"email" bson:"email"`
		Password  string        `json:"password" bson:"password"`
		//School     mgo.DBRef       `json:"school" bson:"school"`
		//PrevSchool []mgo.DBRef     `json:"prev_school" bson:"prev_school"`
		School     bson.ObjectId   `json:"school" bson:"school,omitempty"`
		PrevSchool []bson.ObjectId `json:"prev_school" bson:"prev_school,omitempty"`
		Roles      []Role          `json:"roles" bson:"roles"`
		Status     enums.Status    `json:"status" bson:"status"`
		Access     enums.Access    `json:"access" bson:"access"`
	}
)
