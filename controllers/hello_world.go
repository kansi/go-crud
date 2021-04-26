package controllers

import (
	"net/http"
)

type AppData struct {
	DataId  string `db:"data_id"`
	Message string `db:"message"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
