package web

import (
	"net/http"

	"gopkg.in/labstack/echo.v2"
)

type dbConn interface {
	ExecQuery(string) (map[string]interface{}, error)
}

func New(conn dbConn) http.Handler {
	ech := echo.New()

	return ech
}
