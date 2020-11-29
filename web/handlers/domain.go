package handlers

import (
	"net/http"
	"strconv"

	"github.com/duraki/decipiat/core"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func DomainView(c echo.Context) error {
	log.Infof("Domain view started")
	return c.Render(http.StatusOK, "domain", nil)
}

func SearchDomain(c echo.Context) error {
	// Collect form data
	domain := c.FormValue("domain")
	number, _ := strconv.Atoi(c.FormValue("number"))

	// Collect data from checkboxes
	var types string
	possibleTypes := []string{"all", "bitsquatting", "vowels", "repetition", "omission", "homograph", "hyphenation"}

	if c.FormValue("all") != "" {
		types = "all"
	} else {
		for _, v := range possibleTypes {
			if val := c.FormValue(v); val != "" {
				types += val + "+"
			}
		}
	}

	domains, total := core.GenerateSimilar(domain, number, types)

	return c.Render(http.StatusOK, "domainview", map[string]interface{}{
		"total":   total,
		"domain":  domain,
		"domains": domains,
	})
}
