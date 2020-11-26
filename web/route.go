package web

import (
	_ "github.com/duraki/decipiat/web/api"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/gommon/log"
	_ "github.com/labstack/echo/v4/middleware"
	_ "log"
)

/* sample -- https://github.com/xesina/golang-echo-realworld-example-app/tree/master/router */
func Init() *echo.Echo {
	e := echo.New()

	/*
		v1 := e.Group("/api/v1")
		{
			v1.GET("/", api.Home())
		}
	*/

	return e
}
