package middlewares

import (
	"fmt"
	"github.com/duraki/decipiat/web/session"
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if session.IsUserAuthenticated(c) {
			return next(c)
		} else {
			return c.Render(http.StatusBadRequest, "message", map[string]interface{}{
				"msg": template.HTML(fmt.Sprintf("Please <a href='/login'>Login</a> first to view this page.")),
			})
		}
	}
}
