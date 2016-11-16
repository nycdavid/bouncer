package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var db

func main() {
	port := os.Getenv("PORT")
  db, err := sql.Open("postgres", "user=postgres dbname=bouncer_dev")
  if err != nil {
    log.Fatal(err)
  }
	e := echo.New()
	e.GET("/", PostHandler)

	if err := e.Start(fmt.Sprintf(":%v", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}

func PostHandler(ctx echo.Context) error {
  rows, _ := db.Query("select count(*) from products")
	return ctx.String(http.StatusOK, string(rows))
}
