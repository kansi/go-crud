package main

import (
	"context"
	"go-crud/app"
	"go-crud/controllers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	connString := os.Getenv("POSTGRESQL_URL")
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Println(err)
		return
	}

	repo := app.NewAppRepo(pool)
	c := controllers.NewAppEnv(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", c.HelloWorld)

	http.ListenAndServe(":7000", r)
}
