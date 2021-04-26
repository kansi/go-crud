package main

import (
	"context"
	"go-crud/app"
	"go-crud/controllers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestHelloWorld(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	connString := os.Getenv("POSTGRESQL_URL")
	pool, _ := pgxpool.Connect(context.Background(), connString)
	repo := app.NewAppRepo(pool)
	c := controllers.NewAppEnv(repo)

	http.HandlerFunc(c.HelloWorld).ServeHTTP(rec, req)

	expected := "Hello World!"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
