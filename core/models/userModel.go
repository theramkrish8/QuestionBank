package models

import "gopkg.in/mgo.v2/bson"

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
	}
	//School object
	School struct {
		SchoolID   int64   `json:"school_id" bson:"_id,omitempty"`
		SchoolName string  `json:"School_name" bson:"school_name"`
		Address    Address `json:"address" bson:"address"`
		Contact    Contact `json:"contact" bson:"contact"`
	}
	//Class object
	Class struct {
		ClassID  bson.ObjectId `json:"class_id" bson:"_id,omitempty"`
		Standard string        `json:"standard" bson:"standard"`
		Sections []string      `json:"sections" bson:"sections"`
		Subjects []Subject     `json:"subjects" bson:"subjects"`
	}
	//Subject object
	Subject struct {
		SubjectID   bson.ObjectId `json:"subject_id" bson:"_id,omitempty"`
		SubjectCode string        `json:"subject_code" bson:"subject_code"`
		SubjectName string        `json:"subject_name" bson:"subject_name"`
		Chapters    []Chapter     `json:"chapters" bson:"chapters"`
	}

	//Chapter object
	Chapter struct {
		ChapterID      int64  `json:"chapter_id" bson:"_id,omitempty"`
		ChapterSquence string `json:"chapter_squence" bson:"chapter_sequence"`
		ChapterName    string `json:"chapter_name" bson:"chapter_name"`
	}
	//Address object
	Address struct {
		Street   string `json:"street" bson:"street"`
		Place    string `json:"place" bson:"place"`
		District string `json:"district" bson:"district"`
		State    string `json:"state" bson:"state"`
		PinCode  string `json:"pin_code" bson:"pin_code"`
	}
	//Contact object
	Contact struct {
		Landline []string `json:"land_line" bson:"Land_line"`
		Mobile   []string `json:"mobile" bson:"mobile"`
		Fax      string   `json:"fax" bson:"fax"`
		Email    string   `json:"email" bson:"email"`
		Website  string   `json:"website" bson:"website"`
	}
)
