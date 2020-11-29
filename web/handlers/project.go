package handlers

import (
	_ "encoding/json"
	"fmt"
	"github.com/duraki/decipiat/models"
	"github.com/duraki/decipiat/web/session"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
	"time"
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
	cpvUuid := uuid.New().String()

	// Bind
	prj := &models.Project{
		ID:         bson.NewObjectId(),
		Name:       name,
		InternalId: int,
		UserID:     session.GetUser().ID,
		CpvUuid:    cpvUuid,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

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
	defer db.Close()

	var projects []models.Project

	// Pull specific:
	if err = db.DB(DatabaseName).C(models.CollectionProject).Find(bson.M{"_userid": bson.ObjectIdHex(fmt.Sprintf("%x", string(session.User.ID)))}).All(&projects); err != nil {
		log.Errorf("%s %+v", "Error while retrieve Project List View, User", session.GetUser())
		return c.Render(http.StatusOK, "message", map[string]interface{}{
			"msg": "Your project list is empty.",
		})
	}

	return c.Render(http.StatusOK, "project_list", map[string]interface{}{
		"projects": projects,
		"user":     session.GetUser(),
	})

	return c.JSON(http.StatusOK, projects)
}

func ProjectView(c echo.Context) (err error) {
	cpvUuid := c.Param("cpvUuid")

	// Bind
	project := new(models.Project)

	db := GlobalConfig.DB.Clone()
	defer db.Close()

	if err = db.DB(DatabaseName).C(models.CollectionProject).
		Find(bson.M{"cpvUuid": cpvUuid}).One(project); err != nil {
		if err == mgo.ErrNotFound {
			return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
				"msg": fmt.Sprintf("Project with uuid <%s> does not exists.", cpvUuid),
			})
		}

		if err != nil {
			log.Errorf("Unknown error occured when selecting project id: %s", cpvUuid)
			return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
				"msg": fmt.Sprintf("Unknown error occured ... Can't select project %s", cpvUuid),
			})

		}

		return
	}

	return c.Render(http.StatusUnauthorized, "project_view", map[string]interface{}{
		"project": project,
		"user": session.GetUser()
	})
}
