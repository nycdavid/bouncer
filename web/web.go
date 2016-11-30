package web

import (
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
	fmt.Println("This is the POST")
	var reqB ReqBody
	deco := json.NewDecoder(ctx.Request().Body())
	err := deco.Decode(&reqB)
	fmt.Println(reqB)
	if err != nil {
		return err
	}
	obj, err := dbc.ExecQuery("select * from products")
	if err != nil {
		return err
	}
	json, err := json.Marshal(obj)
	log.Println(json)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, json)
}
