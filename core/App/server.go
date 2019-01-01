package main

import (
	"github.com/gin-gonic/gin"

	"QuestionBank/core/services"
)

func main() {
	router := gin.Default()
	//User handlers
	addUserHandlers(router)
	//School handlers
	addSchoolHandlers(router)
	//Group handlers
	addGroupHandlers(router)
	//Class handlers
	addClassHandlers(router)
	router.Run(":6060")
}

func addUserHandlers(router *gin.Engine) {
	router.GET("/user", services.UsersList)
	router.POST("/user", services.CreateUser)
	router.PUT("/user", services.UpdateUser)
	router.DELETE("/user", services.RemoveUser)
	router.GET("/user/:id", services.GetUser)
}

func addSchoolHandlers(router *gin.Engine) {
	router.GET("/school", services.SchoolList)
	router.POST("/school", services.CreateSchool)
	router.PUT("/school", services.UpdateSchool)
	router.DELETE("/school", services.RemoveSchool)
	router.GET("/school/:id", services.GetSchool)
}

func addGroupHandlers(router *gin.Engine) {
	router.GET("/group", services.GroupList)
	router.GET("/group/:id", services.GetUser)
	router.POST("/group", services.CreateGroup)
	router.PUT("/group/:id1/user/:id2", services.AddUserToGroup)
	router.PUT("/group/:id1/school/:id2", services.AddSchoolToGroup)
	router.DELETE("/group", services.RemoveGroup)
}

func addClassHandlers(router *gin.Engine) {
	router.GET("/class", services.ClassList)
	router.GET("/class/:id", services.GetClass)
	router.POST("/class", services.CreateClass)
	router.PUT("/class/:id", services.UpdateClass)
	router.DELETE("/class", services.RemoveClass)
}
