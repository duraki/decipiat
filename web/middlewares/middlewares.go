package middlewares

import (
	"net/http"

	"github.com/duraki/decipiat/web/session"
	"github.com/labstack/echo"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if session.IsUserAuthenticated(c) {
			return next(c)
		} else {
			return c.Render(http.StatusBadRequest, "message", map[string]interface{}{
				"msg": "Empty Email or Password",
			})
		}
	}
}
