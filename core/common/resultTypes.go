package common

import (
	"gopkg.in/mgo.v2/bson"

	"QuestionBank/core/enums"
	"QuestionBank/core/models"
)

type (
	//UserResult object to return API response
	UserResult struct {
		UserID     bson.ObjectId   `json:"user_id" bson:"_id,omitempty"`
		FirstName  string          `json:"first_name" bson:"first_name"`
		LastName   string          `json:"last_name" bson:"last_name"`
		Gender     string          `json:"gender" bson:"gender"`
		Mobile     string          `json:"mobile" bson:"mobile"`
		Email      string          `json:"email" bson:"email"`
		Password   string          `json:"password" bson:"password"`
		School     models.School   `json:"school" bson:"school,omitempty"`
		PrevSchool []models.School `json:"prev_school" bson:"prev_school,omitempty"`
		Roles      []models.Role   `json:"roles" bson:"roles"`
		Status     enums.Status    `json:"status" bson:"status"`
		Access     enums.Access    `json:"access" bson:"access"`
	}
)
