package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/bmizerany/pq"
	"github.com/labstack/echo/engine/standard"
	"github.com/nycdavid/bouncer/web"
)

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

func main() {
	port := os.Getenv("PORT")
	pgDbo, err := sql.Open("postgres", "user=postgres port=32768 dbname=bouncer_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	pgc := web.PGConn{Dbo: pgDbo}
	web := web.New(pgc)

	e.Run(standard.New(fmt.Sprintf(":%v", port)))
}
