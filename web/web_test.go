package web

import (
	"bytes"
	"encoding/json"
	"testing"

	"gopkg.in/labstack/echo.v2/test"
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
	var b bytes.Buffer
	reqBody := ReqBody{Ids: []int{2}}
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", json.NewEncoder(b).encode(reqBody))
	web := New(MockConn{})
	ctx := web.NewContext(req, rec)
	PostHandler(ctx)

	deco := json.NewDecoder(rec.Body)
	err := deco.Decode(&respBody)
	if err != nil {
		t.Error(err)
	}

	if rec.Status() != 200 {
		t.Errorf("Expected status code to be 200, but got %v", rec.Status())
	}
}
