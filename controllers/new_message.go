package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

func (env *AppEnv) CreateNewMessage(c *gin.Context) {
	var request CreateMessageRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := env.repo.AppData.StoreMessage(request.Message)

	if err != nil {
		c.String(http.StatusInternalServerError, "Some error occured")
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": id,
	})
}
