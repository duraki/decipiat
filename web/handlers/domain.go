package handlers

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/duraki/decipiat/core"
	"net/http"
)

func DomainView(c echo.Context) error {
	log.Infof("Domain view started")
	return c.Render(http.StatusOK, "domain", nil)
}

func SearchDomain(c echo.Context) error {
	domain := c.FormValue("domain")

	domains := core.GenerateSimilar(domain)

	return c.Render(http.StatusOK, "domainview", map[string]interface{}{
		"domain": domain,
		"domains": domains,
	})
}
