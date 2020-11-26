package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.String(http.StatusOK, "API v1")
	}
}
