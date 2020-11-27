package handlers

import (
	"fmt"
	"github.com/duraki/decipiat/models"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
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
	/* hash passwords */
	hash, _ := models.HashPassword(c.FormValue("password"))

	// Bind
	u := &models.User{ID: bson.NewObjectId(), Email: email, Password: hash}

	// Default binding via POST req body.
	// if err = c.Bind(u); err != nil {
	// 	return
	// }

	// Validate
	if u.Email == "" || u.Password == "" {
		// taken from here: https://gitlab.com/ykyuen/golang-echo-template-example/-/tree/master
		return c.Render(http.StatusBadRequest, "message", map[string]interface{}{
			"msg": "Empty Email or Password",
		})
		// return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	db := GlobalConfig.DB.Clone()
	defer db.Close()
	if err = db.DB("decipiat").C("users").Insert(u); err != nil {
		return
	}

	return c.Render(http.StatusCreated, "message", map[string]interface{}{
		"msg": fmt.Sprintf("User registered with email: %s", email),
	})

	//return c.JSON(http.StatusCreated, u)
}

func LoginUser(c echo.Context) (err error) {
	// Read the fields
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Bind
	user := new(models.User)

	// Find user
	db := GlobalConfig.DB.Clone()
	defer db.Close()
	if err = db.DB(DatabaseName).C(models.CollectionUser).
		Find(bson.M{"email": email, "password": password}).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	return c.JSON(http.StatusOK, "Logged in")
}
