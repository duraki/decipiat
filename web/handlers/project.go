package handlers

import (
	_ "encoding/json"
	"fmt"
	"github.com/duraki/decipiat/models"
	"github.com/duraki/decipiat/web/session"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
)

func ProjectCreateView(c echo.Context) error {
	if !session.IsUserAuthenticated(c) {
		return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
			"msg": fmt.Sprintf("Please login again to continue."),
		})
	}
	log.Infof("%s\n", "project create view handler started ...")

	/* set user to session */
	user := session.GetUser()

	return c.Render(http.StatusOK, "project_create", map[string]interface{}{
		"user": user,
	})
}

func ProjectCreate(c echo.Context) (err error) {
	if !session.IsUserAuthenticated(c) {
		return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
			"msg": fmt.Sprintf("Please login again to continue."),
		})
	}

	name := c.FormValue("project_name")
	int := c.FormValue("identifier")

	// Bind
	prj := &models.Project{ID: bson.NewObjectId(), Name: name, InternalId: int, UserID: session.GetUser().ID}

	// Validate
	if prj.Name == "" {
		// taken from here: https://gitlab.com/ykyuen/golang-echo-template-example/-/tree/master
		return c.Render(http.StatusBadRequest, "message", map[string]interface{}{
			"msg": "Empty Project Name. This field is required.",
		})
	}

	// Save project
	db := GlobalConfig.DB.Clone()
	defer db.Close()
	if err = db.DB(DatabaseName).C(models.CollectionProject).Insert(prj); err != nil {
		return
	}

	return c.Render(http.StatusCreated, "message", map[string]interface{}{
		"msg":      template.HTML(fmt.Sprintf("Project [%s] created (w internal_id: %s) #.", prj.Name, prj.InternalId)),
		"template": template.HTML(fmt.Sprintf("%s", "Click <a href='/project/list'>here</a> to List all projects.")),
	})
}

func ProjectListView(c echo.Context) (err error) {
	db := GlobalConfig.DB.Clone()

	var projects []*models.Project

	defer db.Close()
	//bson.M{"userId": session.User.ID}
	// Pull All:
	// if err = db.DB(DatabaseName).C(models.CollectionProject).Find(nil).All(&projects); err != nil {
	// 	log.Errorf("%s %+v", "Error while retrieve Project List View, User", session.GetUser())
	// }

	// Pull Specific:
	query := bson.M{
		"$match": bson.M{
			"userId": session.User.ID,
		},
	}
	if err = db.DB(DatabaseName).C(models.CollectionProject).Find(query).All(&projects); err != nil {
		log.Errorf("%s %+v", "Error while retrieve Project List View, User", session.GetUser())
		return
	}

	return c.JSON(http.StatusOK, projects)

	/**
	prj := &models.Project{}

	if prj, err = db.DB(DatabaseName).C(models.CollectionProject)
		.Find(bson.M{"_userId": session.GetUser().ID}); err != nil {
			return c.JSON(http.StatusOK, fmt.Sprintf("Found total projects %+v", ))
		}
	**/
}
