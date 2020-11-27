package handlers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
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

func RegisterUser(c echo.Context) error {

}
