package web

import (
	"net/http"

	"gopkg.in/labstack/echo.v2"
)

type dbConn interface {
	ExecQuery(string) (map[string]interface{}, error)
}

func New(conn dbConn) *echo.Echo {
	ech := echo.New()
	ech.POST("/", PostHandler)

	return ech
}

func PostHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "foo")
}
