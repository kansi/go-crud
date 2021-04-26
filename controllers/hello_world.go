package controllers

import (
	"github.com/jackc/pgx/v4"
	"net/http"
)

type AppData struct {
	DataId  string `db:"data_id"`
	Message string `db:"message"`
	Email   string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
