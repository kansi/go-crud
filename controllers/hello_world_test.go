package controllers

import (
	"context"
	"go-crud/app"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestHelloWorld(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// test setup
	connString := os.Getenv("POSTGRESQL_URL")
	pool, _ := pgxpool.Connect(context.Background(), connString)
	repo := app.NewAppRepo(pool)
	c := NewAppEnv(repo)

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	// setup router
	router := SetupRouter(c)
	router.ServeHTTP(rec, req)

	expected := "{\"message\":\"Hello World!\"}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
