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
	pgDbo, err := sql.Open("postgres", "user=postgres port=32768 dbname=bouncer_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	pgc := web.PGConn{Dbo: pgDbo}
	web := web.New(pgc)
	web.Run(standard.New(":" + port))
}
