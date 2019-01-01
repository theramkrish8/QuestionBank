package models

type (
	//Contact object
	Contact struct {
		Landline []string `json:"land_line" bson:"Land_line"`
		Mobile   []string `json:"mobile" bson:"mobile"`
		Fax      string   `json:"fax" bson:"fax"`
		Email    string   `json:"email" bson:"email"`
		Website  string   `json:"website" bson:"website"`
	}
)
