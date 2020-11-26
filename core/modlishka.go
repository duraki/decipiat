package core

import (
	"io/ioutil"
	"encoding/json"
	"github.com/duraki/decipiat/models"
)

func StartModlishka(mdlConfig models.ModlishkaConfig) {
	file, _ := json.MarshalIndent(mdlConfig," ", " ")
	_ = ioutil.WriteFile(*mdlConfig.ProxyDomain + ".json", file, 0644)
}
