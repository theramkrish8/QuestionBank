package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	//School object
	School struct {
		SchoolID     bson.ObjectId `json:"school_id" bson:"_id,omitempty"`
		SchoolCode   string        `json:"school_code" bson:"school_code"`
		SchoolName   string        `json:"school_name" bson:"school_name"`
		Address      Address       `json:"address" bson:"address"`
		Contact      Contact       `json:"contact" bson:"contact"`
		GroupID      mgo.DBRef     `json:"group_id" bson:"group_id"`
		PrevGroup    []mgo.DBRef   `json:"prev_group" bson:"prev_group"`
		CreatedTime  string        `json:"create_time" bson:"create_time"`
		ModifiedTIme string        `json:"modified_time" bson:"modified_time"`
	}
)
