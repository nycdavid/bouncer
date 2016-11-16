package main

import (
	"fmt"
	"net/http"
	"os"

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
	return ctx.String(http.StatusOK, "Hello World!")
}
