package controllers

import (
	"fmt"
	"go-crud/app"
	"net/http"
)

// AppEnv holds database Repo
type AppEnv struct {
	repo *app.AppRepo
}

func NewAppEnv(repo *app.AppRepo) *AppEnv {
	return &AppEnv{
		repo: repo,
	}
}

func (env *AppEnv) HelloWorld(w http.ResponseWriter, r *http.Request) {
	data, err := env.repo.FindAppDataByID(1)

	if err != nil {
		fmt.Println("Error", data)
	}

	w.Write([]byte(data.Message))
}
