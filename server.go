package main

import (
	"context"
	"go-crud/app"
	"go-crud/controllers"
	"log"
	"os"

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
	r := controllers.SetupRouter(c)

	r.Run(":7000")
}
