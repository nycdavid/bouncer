package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"gopkg.in/labstack/echo.v2/test"
)

type MockConn struct {
}

func (mc MockConn) ExecQuery(sqlQuery string) (map[string]interface{}, error) {
	if sqlQuery == "select id from products where id in (2)" {
		resp := map[string]interface{}{
			"matchedCount": 1,
			"matchedIds":   []int{1, 2},
		}
		return resp, nil
	} else {
		return nil, errors.New("Incorrect SQL string.")
	}
}

func TestPostReturnsMatchedCount(t *testing.T) {
	var respBody RespBody
	reqBody := ReqBody{Ids: []int{2}}
	b, err := json.Marshal(reqBody)
	if err != nil {
		t.Error(err)
	}
	rec := test.NewResponseRecorder()
	req := test.NewRequest("POST", "/", bytes.NewReader(b))
	web := New(MockConn{})
	ctx := web.NewContext(req, rec)
	PostHandler(ctx)

	deco := json.NewDecoder(rec.Body)
	err = deco.Decode(&respBody)
	if err != nil {
		t.Error(err)
	}

	if rec.Status() != 200 {
		t.Errorf("Expected status code to be 200, but got %v", rec.Status())
	}
	if respBody.MatchedCount != 1 {
		t.Errorf("Expected response MatchedCount to be 1, but got %v", respBody.MatchedCount)
	}
}
