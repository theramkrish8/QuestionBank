package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
		Group     mgo.DBRef     `json:"group" bson:"group"`
		School    mgo.DBRef     `json:"school" bson:"school"`
		Roles     []Role        `json:"roles" bson:"roles"`
		Status    Status        `json:"status" bson:"status"`
		Access    Access        `json:"access" bson:"access"`
	}

	//RoleType enum type
	RoleType int
	//Status enum type
	Status int
	//Access enum type
	Access int
	//Role type
	Role struct {
		Role       RoleType
		Classes    []Class
		AssignedOn string
	}
)

//enum type RoleType
const (
	BoardType       RoleType = 1
	GroupType       RoleType = 2
	SchoolType      RoleType = 3
	SubOrdinateType RoleType = 4
	TeacherType     RoleType = 5
)

//enum type Status
const (
	Active  Status = 1
	Dorment Status = 0
)

//enum tye Access
const (
	None     Access = 0
	ReadOnly Access = 1
	All      Access = 2
)
