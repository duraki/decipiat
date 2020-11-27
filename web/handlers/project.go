package handlers

import (
	_ "encoding/json"
	"fmt"
	_ "github.com/duraki/decipiat/models"
	"github.com/duraki/decipiat/web/session"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ProjectCreateView(c echo.Context) error {
	if !session.IsUserAuthenticated(c) {
		log.Infof("%s\n", "user is not authenticated, trying to access pass the project page")
		return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
			"msg": fmt.Sprintf("Please login again to continue."),
		})
		//return c.Render(http.StatusOK, "userhome", session.GetSessionFromRequest(c).User)
	}
	log.Infof("%s\n", "project create view handler started ...")

	/* set user to session */
	user := session.GetUser()

	return c.Render(http.StatusOK, "project_create", map[string]interface{}{
		"user": user,
	})
}

/**
func ProjectCreate(c echo.Context) error {
	//return c.String(http.StatusOK, "Project created")
	project := models.Project{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&project)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	log.Printf("new project created .. in %#v", project)
	return c.String(http.StatusOK, "Project created")
}
**/

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
