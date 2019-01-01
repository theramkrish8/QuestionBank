package models

type (
	//Subject object
	Subject struct {
		SubjectCode string    `json:"subject_code" bson:"subject_code"`
		SubjectName string    `json:"subject_name" bson:"subject_name"`
		Chapters    []Chapter `json:"chapters" bson:"chapters"`
	}
)
