package handlers

import (
	"github.com/duraki/decipiat/web/session"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func Homepage(c echo.Context) error {
	log.Println("Homepage handler started ... exec øOK")
	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"user": session.GetUser(),
	})
}
