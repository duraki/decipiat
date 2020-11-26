package handlers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func Status(c echo.Context) error {
	log.Println("Status handler started ... exec øOK")
	return c.JSON(http.StatusOK, "Status OK")
}
