package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	//Tags struct
	Tags struct {
		TagID       bson.ObjectId `json:"tag_id" bson:"_id omitempty"`
		Description string        `json:"descriptio" bson:"description"`
	}
)
