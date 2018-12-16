package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Group object
type Group struct {
	GroupID      bson.ObjectId `json:"group_id" bson:"_id,omitempty"`
	GroupName    string        `json:"group_name" bson:"group_name"`
	GroupCode    string        `json:"group_code" bson:"group_code"`
	Board        string        `json:"board" bson:"board"`
	Users        []mgo.DBRef   `json:"users" bson:"users"`
	Schools      []mgo.DBRef   `json:"schools" bson:"schools"`
	Classes      []Class       `json:"classes" bson:"classes"`
	Questions    []Question    `json:"questions" bson:"questions"`
	AuxQuestions []Question    `json:"aux_questions" bson:"aux_questions"`
}
