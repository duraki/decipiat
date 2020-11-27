package handlers

import (
	"fmt"
	"github.com/duraki/decipiat/models"
	"github.com/duraki/decipiat/web/session"
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

	if session.IsUserAuthenticated(c) {
		return c.Redirect(http.StatusTemporaryRedirect, "/project/new")
		/**
		return c.Render(http.StatusOK, "message", map[string]interface{}{
			"msg": fmt.Sprintf("Please login again to continue."),
		})
		**/
	}

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
	if err = db.DB(DatabaseName).C(models.CollectionUser).Insert(u); err != nil {
		return
	}

	return c.Render(http.StatusCreated, "message", map[string]interface{}{
		"msg": fmt.Sprintf("User registered with email: %s", email),
	})

	//return c.JSON(http.StatusCreated, u)
}

func LogoutUser(c echo.Context) (err error) {
	// Kill sessions
	session, _ := session.Store.Get(c.Request(), session.SessionTokenName)
	session.Values["authenticated"] = false
	session.Save(c.Request(), c.Response())

	return c.Render(http.StatusOK, "message", map[string]interface{}{
		"msg": fmt.Sprintf("Thank you for operating on decipiat. Come again for new logins."),
	})
}

func LoginUser(c echo.Context) (err error) {
	// Read the fields
	email := c.FormValue("email")

	// Bind
	user := new(models.User)

	// Find user
	db := GlobalConfig.DB.Clone()
	defer db.Close()
	if err = db.DB(DatabaseName).C(models.CollectionUser).
		Find(bson.M{"email": email}).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
				"msg": fmt.Sprintf("User with email <%s> does not exists.", email),
			})
		}
		if err != nil {
			return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
				"msg": fmt.Sprintf("Unknown error occured ..."),
			})
			log.Errorf("Unknown error occured when logging in: %s", email)
		}
		return
	}

	session, _ := session.Store.Get(c.Request(), session.SessionTokenName)
	if models.HashCompare(c.FormValue("password"), user.Password) {
		session.Values["authenticated"] = true
		session.Values["email"] = email
		session.Save(c.Request(), c.Response())

		return c.Render(http.StatusOK, "userhome", map[string]interface{}{
			"email": user.Email,
		})
		/**
		return c.Render(http.StatusOK, "message", map[string]interface{}{
			"msg": fmt.Sprintf("User is logged-in."),
		})
		*/
	}

	return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
		"msg": fmt.Sprintf("Incorrect authentication details."),
	})
}
