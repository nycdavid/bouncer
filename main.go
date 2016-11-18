package main

import (
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

func main() {
	port := os.Getenv("PORT")
	e := echo.New()
	e.POST("/", PostHandler)

	if err := e.Start(fmt.Sprintf(":%v", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}

func PostHandler(ctx echo.Context) error {
	db, err := sql.Open("postgres", "user=postgres port=32768 dbname=bouncer_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// rows, _ := db.Query("select * from products")
	// defer rows.Close()
	// var id int
	// var name string
	// for rows.Next() {
	// 	err = rows.Scan(&id, &name)
	// }
	var rb ReqBody
	d := json.NewDecoder(ctx.Request().Body)
	d.Decode(&rb)
	fmt.Println(rb.Ids)

	return ctx.String(http.StatusOK, "Hello world!")
}
