package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/bmizerany/pq"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type ReqBody struct {
	Ids []int `json:"ids"`
}

type RespBody struct {
	MatchedCount int   `json:"matchedCount"`
	MatchedIds   []int `json:"matchedIds"`
}

type QuerySender interface {
	QuerySend(string) (map[string]interface{}, error)
}

// PGSender methods
type PGSender struct {
	Dbo *sql.DB
}

func (pgs PGSender) QuerySend(sql string) (map[string]interface{}, error) {
	var aggInfo map[string]interface{}
	var matchedIds []int

	rows, err := pgs.Dbo.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		matchedIds = append(matchedIds, id)
	}

	aggInfo["matchedCount"] = len(matchedIds)
	aggInfo["matchedIds"] = matchedIds
	return aggInfo, nil
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

var db *sql.DB
var pgSender PGSender
var err error

func main() {
	port := os.Getenv("PORT")
	e := echo.New()

	// DB connection
	db, err = sql.Open("postgres", "user=postgres port=32768 dbname=bouncer_dev sslmode=disable")
	pgSender = PGSender{
		Dbo: db,
	}
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	e.POST("/", PostHandler)
	e.Run(standard.New(fmt.Sprintf(":%v", port)))
}

func PostHandler(ctx echo.Context) error {
	var rb ReqBody

	d := json.NewDecoder(ctx.Request().Body())
	d.Decode(&rb)
	sqlString := ConstructQuery(rb.Ids)
	aggInfo, _ := pgSender.QuerySend(sqlString)

	respBody := RespBody{
		MatchedCount: aggInfo["matchedCount"].(int),
		MatchedIds:   aggInfo["matchedIds"].([]int),
	}

	return ctx.JSON(http.StatusOK, &respBody)
}
