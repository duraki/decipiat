package core

import (
	"encoding/json"
	"encoding/hex"
	"net/http"
	"io/ioutil"
	"github.com/duraki/decipiat/models"

	log "github.com/sirupsen/logrus"
)

type Domain struct {
	Name string `json:"domain"`
	DomainAsHex string `json:"domain_as_hexadecimal"`
	Fuzz_URL string `json:"fuzz_url"`
	Fuzzer string `json:"fuzzer"`
	ParkedScoreURL string `json:"parked_score_url"`
	ResolveIPURL string `json:"resolve_ip_url"`
}

type Domains struct {
	Name string `json:"domain"`
	DomainAsHex string `json:"domain_as_hexadecimal"`
	FuzzyDomains []Domain `json:"fuzzy_domains"`
	ParkedScoreURL string `json:"parked_score_url"`
	ResolveIPURL string `json:"resolve_ip_url"`
	URL string `json:"url"`
}

func GenerateSimilar(domain string) *models.AllDomains {
	domain = stringToHex(domain)

	resp, err := http.Get("https://dnstwister.report/api/fuzz/696e6669676f2e6872")
	if err != nil {
		log.Infof("GenerateSimilar: Could not fetch domain\n")
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var wholeDomainFetch Domains
	_ = json.Unmarshal(body, &wholeDomainFetch)

	var domains models.AllDomains

	for _, v := range wholeDomainFetch.FuzzyDomains {
		domains.All = append(domains.All, v.Name)
	}

	domains.Available = checkAvailable(wholeDomainFetch.FuzzyDomains)


	return &domains
}

func checkAvailable(domains []Domain) []string {
	var availableDomains []string

	type availableDomain struct {
		Domain string `json:"domain"`
		IP string `json:"ip"`
	}

	for _, v := range domains {
		var available availableDomain
		resp, err := http.Get(v.ResolveIPURL)
		if err != nil {
			log.Infof("checkAvailable: could not fetch ip data about domain\n")
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		_ = json.Unmarshal(body, &available)

		if available.IP != "" {
			availableDomains = append(availableDomains, available.Domain)
		}
	}

	return availableDomains
}

func stringToHex(word string) string {
	str := []byte(word)
	return hex.EncodeToString(str)
}
