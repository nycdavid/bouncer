package web

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/labstack/echo.v2"
)

var dbc dbConn

type dbConn interface {
	ExecQuery(string) (map[string]interface{}, error)
}

type PGConn struct {
	Dbo *sql.DB
}

type ReqBody struct {
	Ids []int `json:"ids"`
}

func ConstructQuery(reqIds []int) string {
	var sqlBuffer bytes.Buffer
	sqlBuffer.WriteString("select id from products where id in (")
	for i := 0; i < len(reqIds); i++ {
		id := reqIds[i]
		if i == len(reqIds)-1 {
			sqlBuffer.WriteString(fmt.Sprintf("%v", id))
		} else {
			sqlBuffer.WriteString(fmt.Sprintf("%v, ", id))
		}
	}
	sqlBuffer.WriteString(")")

	return sqlBuffer.String()
}

func (pgc PGConn) ExecQuery(sqlString string) (map[string]interface{}, error) {
	var matchedIds []int

	rows, err := pgc.Dbo.Query(sqlString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		matchedIds = append(matchedIds, id)
	}
	return map[string]interface{}{
		"matchedCount": len(matchedIds),
		"matchedIds":   matchedIds,
	}, nil
}

type RespBody struct {
	MatchedCount int   `json:"matchedCount"`
	MatchedIds   []int `json:"matchedIds"`
}

func New(conn dbConn) *echo.Echo {
	ech := echo.New()
	ech.POST("/", PostHandler)
	dbc = conn

	return ech
}

func PostHandler(ctx echo.Context) error {
	var reqB ReqBody
	deco := json.NewDecoder(ctx.Request().Body())
	err := deco.Decode(&reqB)
	if err != nil {
		log.Fatal(err)
	}
	sqlString := ConstructQuery(reqB.Ids)
	obj, err := dbc.ExecQuery(sqlString)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, obj)
}
