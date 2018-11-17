package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewUserController(getSession())
	router := gin.Default()
	router.GET("/user", uc.UsersList)
	router.POST("/user", uc.CreateUser)
	router.PUT("/user", uc.UpdateUser)
	router.DELETE("/user", uc.RemoveUser)
	router.GET("/user/{id}", uc.GetUser)

	router.Run(":6060")
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
