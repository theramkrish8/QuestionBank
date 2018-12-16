package models

import "gopkg.in/mgo.v2/bson"

type (
	//Class object
	Class struct {
		ClassID  bson.ObjectId `json:"class_id" bson:"_id,omitempty"`
		Standard string        `json:"standard" bson:"standard"`
		Sections []string      `json:"sections" bson:"sections"`
		Subjects []Subject     `json:"subjects" bson:"subjects"`
	}

	//Chapter object
	Chapter struct {
		ChapterID      int64  `json:"chapter_id" bson:"_id,omitempty"`
		ChapterSquence string `json:"chapter_squence" bson:"chapter_sequence"`
		ChapterName    string `json:"chapter_name" bson:"chapter_name"`
	}

	//Subject object
	Subject struct {
		SubjectID   bson.ObjectId `json:"subject_id" bson:"_id,omitempty"`
		SubjectCode string        `json:"subject_code" bson:"subject_code"`
		SubjectName string        `json:"subject_name" bson:"subject_name"`
		Chapters    []Chapter     `json:"chapters" bson:"chapters"`
	}
)
