package models

type (
	//Question object
	Question struct {
	}
	//Severity enum type
	Severity int
)

//enum type Severity
const (
	Easy     Severity = 1
	Moderate Severity = 2
	Hard     Severity = 3
)
