package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	Latitude string `bson:"latitude" json:"latitude"`
	Longitue string `bson:"longitude" json:"longitude"`
}
type Address struct {
	Street   string
	City     string
	State    string
	Pincode  int
	location Location
}
type STATUS int

const (
	DORMENT STATUS = 0
	ACTIVE  STATUS = 1
)

type ACCESS int

const (
	NONE      ACCESS = 0
	READ_ONLY ACCESS = 1
	ALL       ACCESS = 2
)

type ROLE int

const (
	BOARD       ROLE = 1
	GROUP       ROLE = 2
	SCHOOL      ROLE = 3
	SUBORDINATE ROLE = 4
	TEACHER     ROLE = 5
)

type SEVERITY int

const (
	EASY     SEVERITY = 0
	MODERATE SEVERITY = 1
	HARD     SEVERITY = 2
)

type QUESTION_APPROVAL_STATUS int

const (
	WAITING  QUESTION_APPROVAL_STATUS = 0
	APPROVED QUESTION_APPROVAL_STATUS = 1
	REJECTED QUESTION_APPROVAL_STATUS = 2
)

type Class struct {
	Classname   string
	ClassNumber int
	Section     int
}
type Contact struct {
	Landline []string
	Mobile   []string
	Fax      string
	Email    string
	Website  string
}
type School struct {
	SchoolId      bson.ObjectId `bson:"_schoolid" json:"schoolid"`
	SchoolName    string        `bson:"schoolname" json:"schoolname"`
	SchoolCode    string        `bson:"schoolcode" json:"schoolcode"`
	SchoolAddress Address
	SchoolContact Contact
	Status        STATUS
}
type Board struct {
	BoardId   int
	BoardName string
	State     string
}
type Group struct {
	GroupId   bson.ObjectId `bson:"_id" json:"id"`
	GroupName string        `bson:"groupname" json:"groupname"`
	GroupCode string        `bson:"groupcode" json:"groupcode"`
	Board     *Board
}
type Teacher struct {
	UserId    bson.ObjectId `bson:"_id" json:"id"`
	FirstName string        `bson:"firstname" json:"firstname"`
	LastName  string        `bson:"lastname" json:"lastname"`
	Mobile    int64         `bson:"mobile" json:"mobile"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"password"`
	Group     Group
	School    School
	Classes   []Class
	Roles     []ROLE
	Status    STATUS
	Access    ACCESS
}

type ChapterType struct {
	ChapterId       int64
	ChapterSequence string
	ChapterName     string
}
type TagType struct {
	TagId       int64
	Description string
}
type Question struct {
	QuestionId          int64
	Description         string
	Chapter             ChapterType
	Tags                []TagType
	Severity            SEVERITY
	Marks               int
	ApprovalStatus      QUESTION_APPROVAL_STATUS
	ApprovalDescription string
	Status              STATUS
}

func AllTeachersEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindTeacherEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateTeacherEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func UpdateTeacherEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteTeacherEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/teacher", AllTeachersEndPoint).Methods("GET")
	r.HandleFunc("/teacher", CreateTeacherEndPoint).Methods("POST")
	r.HandleFunc("/teacher", UpdateTeacherEndPoint).Methods("PUT")
	r.HandleFunc("/teacher", DeleteTeacherEndPoint).Methods("DELETE")
	r.HandleFunc("/teacher/{id}", FindTeacherEndpoint).Methods("GET")
	if err := http.ListenAndServe(":6060", r); err != nil {
		log.Fatal(err)
	}
}
