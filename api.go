package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
