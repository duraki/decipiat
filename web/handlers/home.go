package handlers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func Homepage(c echo.Context) error {
	log.Println("Homepage handler started ... exec øOK")
	return c.Render(http.StatusOK, "homepage", nil)
}
