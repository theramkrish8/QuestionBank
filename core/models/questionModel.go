package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/enums"
)

type (
	//Question object
	Question struct {
		QuestionID          bson.ObjectId  `json:"question_id" bson:"_id,omitempty"`
		Description         string         `json:"description" bson:"description"`
		Chapter             mgo.DBRef      `json:"chapter" bson:"chapter"`
		Subject             mgo.DBRef      `json:"subject" bson:"subject"`
		Tags                []Tags         `json:"tags" bson:"tags"`
		Severity            enums.Severity `json:"severity" bson:"severity"`
		Marks               int16          `json:"marks" bson:"marks"`
		ApprovalStatus      enums.Approval `json:"approval_status" bson:"approval_status"`
		ApprovalDescription string         `json:"approval_description" bson:"approval_description"`
		Status              enums.Status   `json:"status" bson:"status"`
	}
)
