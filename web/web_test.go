package web

import (
	"encoding/json"
	"testing"

	"github.com/labstack/echo/test"
)

type MockConn struct {
}

func (mc MockConn) ExecQuery(string) (map[string]interface{}, error) {
	resp := map[string]interface{}{
		"MatchedCount": 1,
		"MatchedIds":   []int{1, 2},
	}
	return resp, nil
}

func TestPostReturnsMatchedCount(t *testing.T) {
	var respBody RespBody
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", nil)
	web := New(MockConn{})
	web.Ech.NewContext(req, rec)
	deco := json.NewDecoder(rec.Body)
	err := deco.Decode(&respBody)
	if err != nil {
		t.Error(err)
	}

	if rec.Status() != 200 {
		t.Errorf("Expected status code to be 200, but got %v", rec.Status())
	}
}
