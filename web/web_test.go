package web

import (
	"testing"

	"github.com/labstack/echo/test"
)

type MockConn struct {
}

func (mc MockConn) ExecQuery(string) (map[string]interface{}, error) {
	resp := map[string]interface{}{
		"foo": "bar",
	}
	return resp, nil
}

func TestPostReturnsMatchedCount(t *testing.T) {
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", nil)

	web := New(MockConn{})
	web.NewContext(req, rec)

	if rec.Code != 200 {
		t.Errorf("Expected status code to be 200, but got %v", rec.Code)
	}
}
