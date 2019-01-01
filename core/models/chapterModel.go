package models

type (
	//Chapter object
	Chapter struct {
		ChapterSquence string `json:"chapter_squence" bson:"chapter_sequence"`
		ChapterName    string `json:"chapter_name" bson:"chapter_name"`
	}
)
