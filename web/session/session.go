package session

import (
	_ "fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	_ "net/http"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("killd-mystic-k3y")
	Store = sessions.NewCookieStore(key)
)

const (
	SessionTokenName = "auth_token"
)

func IsUserAuthenticated(c echo.Context) bool {
	session, _ := Store.Get(c.Request(), SessionTokenName)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		log.Infof("%s\n", "user is not authenticated, trying to access pass the login page")
		return false
		/**
		return c.Render(http.StatusUnauthorized, "message", map[string]interface{}{
			"msg": fmt.Sprintf("Please login again to continue."),
		})
		**/
	}

	return true
}
