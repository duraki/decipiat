package session

import (
	_ "fmt"
	"github.com/duraki/decipiat/models"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	_ "net/http"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("killd-mystic-k3y")
	Store = sessions.NewCookieStore(key)
	User  *models.User
)

const (
	SessionTokenName = "auth_token"
)

func GetEmail() string {
	return User.Email
}

func GetUser() *models.User {
	return User
}

func IsUserAuthenticated(c echo.Context) bool {
	return User != nil

	session, _ := Store.Get(c.Request(), SessionTokenName)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Infof("%s\n", "user is not authenticated, trying to access pass the login page")
		return false
	}

	return true
}

func GetSessionFromRequest(req echo.Context) (sess *sessions.Session) {
	sess, _ = Store.Get(req.Request(), SessionTokenName)
	return sess
}
