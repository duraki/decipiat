package handlers

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func RegisterUserView(c echo.Context) error {
	log.Println("Register user handler started ....")
	return c.Render(http.StatusOK, "register", nil)
}
