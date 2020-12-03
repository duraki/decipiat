package project

import (
	_ "encoding/json"
	"fmt"
	"github.com/duraki/decipiat/models"
	m "github.com/duraki/decipiat/models/prj"
	h "github.com/duraki/decipiat/web/handlers"
	"github.com/duraki/decipiat/web/session"
	_ "github.com/google/uuid"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	_ "html/template"
	"net/http"
	_ "time"
)

func TargetElementariesView(c echo.Context) (err error) {
	cpvUuid := c.Param("cpvUuid")

	// Bind
	project := new(models.Project)

	db := h.GlobalConfig.DB.Clone()
	defer db.Close()

	if err = db.DB(h.DatabaseName).C(models.CollectionProject).
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

	return c.Render(http.StatusOK, "config_elementaries", map[string]interface{}{
		"project": project,
		"user":    session.GetUser(),
	})
}

func TargetElementaries(c echo.Context) (err error) {
	cpvUuid := c.Param("cpvUuid")
	targetUri := c.FormValue("target")
	redirectUri := c.FormValue("redirect")

	// Bind
	target := &m.TargetElementaries{
		ID:          bson.NewObjectId(),
		CpvUuid:     cpvUuid,
		TargetURI:   targetUri,
		RedirectURI: redirectUri,
	}
	// target.CreatedAt, target.UpdatedAt = time.Now(), time.Now()

	db := h.GlobalConfig.DB.Clone()
	defer db.Close()
	if err = db.DB(h.DatabaseName).C(m.CollectionTargetElementaries).Insert(target); err != nil {
		return
	}

	return c.Render(http.StatusCreated, "message", map[string]interface{}{
		"msg": fmt.Sprintf("Target elementaries updated successfully."),
	})
}
