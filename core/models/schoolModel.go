package models

import "gopkg.in/mgo.v2/bson"

type (
	//School object
	School struct {
		SchoolID     bson.ObjectId `json:"school_id" bson:"_id,omitempty"`
		SchoolName   string        `json:"school_name" bson:"school_name"`
		Address      Address       `json:"address" bson:"address"`
		Contact      Contact       `json:"contact" bson:"contact"`
		CreatedTime  string        `json:"create_time" bson:"create_time"`
		ModifiedTIme string        `json:"modified_time" bson:"modified_time"`
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
