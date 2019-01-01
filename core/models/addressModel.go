package models

type (
	//Address object
	Address struct {
		Street   string `json:"street" bson:"street"`
		Place    string `json:"place" bson:"place"`
		District string `json:"district" bson:"district"`
		State    string `json:"state" bson:"state"`
		PinCode  string `json:"pin_code" bson:"pin_code"`
	}
)
