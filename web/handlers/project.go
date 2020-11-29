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
	if err = db.DB(DatabaseName).C(models.CollectionProject).Find(bson.M{"userId": session.User.ID}).All(&projects); err != nil {
		fmt.Errorf("%s %+v", "Error while retrieve Project List View, User", session.GetUser())
	}

	// err := db.C("client").Find(nil).All(&results)
	// if err != nil {
	// 	// TODO: Do something about the error
	// } else {
	// 	fmt.Println("Results All: ", results)
	// }

	// projects := []*models.Project{}

	// if err = db.DB(DatabaseName).C(models.CollectionProject).
	// 	.Find(bson.M({}).
	// 	All(&projects); err != nil {
	// }

	// if err = db.DB(DatabaseName).C(models.CollectionProject).
	// 	Find(bson.M{"userId": session.User.ID}).
	// 	All(&projects); err != nil {
	// 	return
	// }

	return c.JSON(http.StatusOK, projects)

	// // Retrieve posts from database
	// posts := []*model.Post{}
	// db := h.DB.Clone()
	// if err = db.DB("twitter").C("posts").
	// 	Find(bson.M{"to": userID}).
	// 	Skip((page - 1) * limit).
	// 	Limit(limit).
	// 	All(&posts); err != nil {
	// 	return
	// }

	/**
	prj := &models.Project{}

	if prj, err = db.DB(DatabaseName).C(models.CollectionProject)
		.Find(bson.M{"_userId": session.GetUser().ID}); err != nil {
			return c.JSON(http.StatusOK, fmt.Sprintf("Found total projects %+v", ))
		}
	**/
}

/**
if err = db.DB("twitter").C("users").
	Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
	if err == mgo.ErrNotFound {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	}
	return
}
*/

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

*/
