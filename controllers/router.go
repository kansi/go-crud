package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(c *AppEnv) *gin.Engine {
	r := gin.Default()

	r.GET("/", c.HelloWorld)
	r.POST("/message", c.CreateNewMessage)

	return r
}
