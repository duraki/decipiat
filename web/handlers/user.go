package handlers

import (
	"github.com/duraki/decipiat/models"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func RegisterUserView(c echo.Context) error {
	log.Infof("%s\n", "register user view handler started ...")
	return c.Render(http.StatusOK, "register", nil)
}

func LoginUserView(c echo.Context) error {
	log.Infof("%s\n", "login user view handler started ...")
	return c.Render(http.StatusOK, "login", nil)
}

/*
func RegisterUser(c echo.Context) error {
	return c.String(http.StatusOK, "Register me")
}
*/

func RegisterUser(c echo.Context) (err error) {
	// Read the fields
	email := c.FormValue("email")
	//password := []byte(c.FormValue("password"))

	/* hash passwords */
	hash, _ := models.HashPassword(c.FormValue("password"))

	// Bind
	u := &models.User{ID: bson.NewObjectId(), Email: email, Password: hash}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	db := GlobalConfig.DB.Clone()
	defer db.Close()
	if err = db.DB("decipiat").C("users").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}
