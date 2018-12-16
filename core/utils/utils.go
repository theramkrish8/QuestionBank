package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MessageTypeDefault(msg string, c *gin.Context) {
	content := gin.H{
		"status": "200",
		"result": msg,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(201, content)
}

func CheckAndDisplay(err error, msg string, status string, c *gin.Context) {
	if err != nil {
		log.Fatalln(msg, err)
		content := gin.H{
			"status": status,
			"result": msg,
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(200, content)
		panic(err)
	}
}

func DisplayError(msg string, status string, c *gin.Context) {
	content := gin.H{
		"status": status,
		"result": msg,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(200, content)
}
