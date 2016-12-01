package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/bmizerany/pq"
	"github.com/nycdavid/bouncer/web"
	"gopkg.in/labstack/echo.v2/engine/standard"
)

func main() {
	port := os.Getenv("PORT")
	var dbUrl string
	if os.Getenv("DATABASE_URL") == "" {
		dbUrl = "user=postgres port=32768 dbname=bouncer_dev sslmode=disable"
	} else {
		dbUrl = os.Getenv("DATABASE_URL")
	}
	pgDbo, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	pgc := web.PGConn{Dbo: pgDbo}
	web := web.New(pgc)
	web.Run(standard.New(":" + port))
}
