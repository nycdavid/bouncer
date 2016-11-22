package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func TestPostReturnsMatchedCount(t *testing.T) {
	reqBody := `{ "ids": [1, 3, 4] }`
	e := echo.New()
	req, err := http.NewRequest("POST", "/", strings.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	rec := httptest.NewRecorder()
	ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	PostHandler(ctx)

	if rec.Code != 200 {
		t.Errorf("Expected status code to be 200, but got %v", rec.Code)
	}
}
