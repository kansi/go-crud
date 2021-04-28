package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-crud/app"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestCreateNewMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// test setup
	connString := os.Getenv("POSTGRESQL_URL")
	pool, _ := pgxpool.Connect(context.Background(), connString)
	repo := app.NewAppRepo(pool)
	c := NewAppEnv(repo)
	// setup router
	router := SetupRouter(c)

	t.Run("creates new message in DB", func(t *testing.T) {
		rec := httptest.NewRecorder()
		var jsonStr = []byte(`{"message":"A new message"}`)
		req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rec, req)

		expected := "{\"message\":\"inserted\"}"
		var responseJSON map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJSON)
		fmt.Println(responseJSON)

		if responseJSON["message"] == nil {
			t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
		}
	})
}
