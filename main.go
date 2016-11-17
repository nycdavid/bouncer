package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/bmizerany/pq"
	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/", PostHandler)

	if err := e.Start(fmt.Sprintf(":%v", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}

func PostHandler(ctx echo.Context) error {
	db, err := sql.Open("postgres", "user=postgres dbname=bouncer_dev")
	if err != nil {
		log.Fatal(err)
	}
	rows, _ := db.Query("select count(*) from products")
	fmt.Println(rows)
	return ctx.String(http.StatusOK, "Hello world!")
}
