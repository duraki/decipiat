package handlers

import (
	"encoding/json"
	"github.com/duraki/decipiat/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

/*
func ProjectCreateView(c echo.Context) error {

}
*/

func ProjectCreate(c echo.Context) error {
	return c.String(http.StatusOK, "Project created")
	/*
		project := models.Project{}
		defer c.Request().Body.Close()

		err := json.NewDecoder(c.Request().Body).Decode(&project)
		if err != nil {
			log.Fatalf("Failed reading the request body %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
		}

		log.Printf("new project created .. in %#v", project)
		return c.String(http.StatusOK, "Project created")
	*/
}

/*

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

*/
