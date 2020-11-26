package web

import (
	"encoding/json"
	"fmt"
	_ "github.com/duraki/decipiat/web/api"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

/* sample -- https://github.com/xesina/golang-echo-realworld-example-app/tree/master/router */
func Init() *echo.Echo {
	e := echo.New()
	return e
}

func GetHome(c echo.Context) error {
	debug := c.QueryParam("debug")
	return c.String(http.StatusOK, fmt.Sprintf("is debug: %s\n", debug))
}

func GetProject(c echo.Context) error {
	name := c.QueryParam("name")
	dataType := c.QueryParam("data")
	projectType := c.QueryParam("type")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your project ID: %s\n", name))
	} else if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"projectId": name,
			"type":      projectType,
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "please specify the data type as String or JSON",
		})
	}
}

func AddProject(c echo.Context) error {
	type Project struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	project := Project{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&project)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	log.Printf("This is your project %#v", project)
	return c.String(http.StatusOK, "We got your project !!!")

}
