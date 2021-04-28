package controllers

import (
	"fmt"
	"go-crud/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppEnv holds database Repo
type AppEnv struct {
	repo *app.AppRepo
}

func NewAppEnv(appRepo *app.AppRepo) *AppEnv {
	return &AppEnv{
		repo: appRepo,
	}
}

func (env *AppEnv) HelloWorld(c *gin.Context) {
	data, err := env.repo.AppData.FindAppDataByID(1)

	if err != nil {
		fmt.Println("Error", data)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data.Message,
	})
}
