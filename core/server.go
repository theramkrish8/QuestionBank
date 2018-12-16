package main

import (
	"github.com/gin-gonic/gin"

	"QuestionBank/core/services"
)

func main() {
	router := gin.Default()
	//User handlers
	router.GET("/user", services.UsersList)
	router.POST("/user", services.CreateUser)
	router.PUT("/user", services.UpdateUser)
	router.DELETE("/user", services.RemoveUser)
	router.GET("/user/:id", services.GetUser)
	//School handlers
	router.GET("/school", services.SchoolList)
	router.POST("/school", services.CreateSchool)
	router.PUT("/school", services.UpdateSchool)
	router.DELETE("/school", services.RemoveSchool)
	router.GET("/school/:id", services.GetSchool)
	//Group handlers
	router.GET("/group", services.GroupList)
	router.GET("/group/:id", services.GetUser)
	router.POST("/group", services.CreateGroup)
	router.PUT("/group/:id1/user/:id2", services.AddUserToGroup)
	router.PUT("/group/:id1/school/:id2", services.AddSchoolToGroup)
	//router.PUT("/group/{id}/classs", services.AddClassToGroup)
	router.DELETE("/group", services.RemoveGroup)

	router.Run(":6060")
}
