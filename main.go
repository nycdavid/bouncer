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

func (pgs PGSender) QuerySend(string) (map[string]interface{}, error) {
  // rows, err := db.Query(sqlBuffer.String())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
  //
	// for rows.Next() {
	// 	var id int
	// 	err = rows.Scan(&id)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	matchedIds = append(matchedIds, id)
	// }
}

func ConstructQuery([]int) string {
	var sqlBuffer bytes.Buffer
	sqlBuffer.WriteString("select id from products where id in (")
	for i := 0; i < len(rb.Ids); i++ {
		id := rb.Ids[i]
		if i == len(rb.Ids)-1 {
			sqlBuffer.WriteString(fmt.Sprintf("%v", id))
		} else {
			sqlBuffer.WriteString(fmt.Sprintf("%v, ", id))
		}
	}
	sqlBuffer.WriteString(")")

	return sqlBuffer.String(), nil
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
    Dbo: db
  }
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	e.POST("/", PostHandler)
	if err := e.Start(fmt.Sprintf(":%v", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}

func PostHandler(ctx echo.Context) error {
	var rb ReqBody
	var matchedIds []int

	d := json.NewDecoder(ctx.Request().Body)
	d.Decode(&rb)
	sqlString := constructQuery(rb.Ids)
  aggInfo, _ := pgSender.QuerySend(sqlString)

	respBody := RespBody{
		MatchedCount: aggInfo["matchedCount"],
		MatchedIds:   aggInfo["matchedIds"],
	}

	return ctx.JSON(http.StatusOK, &respBody)
}
