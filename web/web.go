package web

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/labstack/echo.v2"
)

type dbConn interface {
	ExecQuery(string) (map[string]interface{}, error)
}

type PGConn struct {
	Dbo *sql.DB
}

type ReqBody struct {
	Ids []int `json:"ids"`
}

func (pgc PGConn) New() PGConn {
	dbo, err := sql.Open("postgres", "user=postgres port=32768 dbname=bouncer_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	pgc.Dbo = dbo
	return pgc
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

type Web struct {
	Ech  *echo.Echo
	Conn dbConn
}

func New(conn dbConn) Web {
	ech := echo.New()
	ech.POST("/", PostHandler)

	return Web{Ech: ech, Conn: conn}
}

func PostHandler(ctx echo.Context) error {
	var reqB ReqBody
	deco := json.NewDecoder(ctx.Request().Body())
	err := deco.Decode(&reqB)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusOK, "Foo")
}
