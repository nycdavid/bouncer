package web

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type MockConn struct {
}

func (mc MockConn) ExecQuery(string) (map[string]interface{}, error) {
  // some dummy code
}

func TestPostReturnsMatchedCount(t *testing.T) {
	reqBody := `{ "ids": [1, 3, 4] }`
  web := New(MockConn{})
	req, err := http.NewRequest("POST", "/", strings.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	rec := httptest.NewRecorder()
	ctx := web.NewContext(standard.NewRequest(req, web.Logger()), standard.NewResponse(rec, web.Logger()))
	PostHandler(ctx)

	if rec.Code != 200 {
		t.Errorf("Expected status code to be 200, but got %v", rec.Code)
	}
}
