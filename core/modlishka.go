package core

import (
	"encoding/json"
	"github.com/duraki/decipiat/models"
	"io/ioutil"
	"os/exec"
)

// StartModlishka will create config file for Modlishka based on params provided and will start Modlishka
func StartModlishka(mdlConfig models.ModlishkaConfig) {
	file, _ := json.MarshalIndent(mdlConfig, " ", " ")
	_ = ioutil.WriteFile(*mdlConfig.ProxyDomain+".json", file, 0644)

	// modlishkaPath := "~/go/src/github.com/drk1wi/Modlishka/dist/proxy"

	cmd := exec.Command("/Users/daemon1/go/src/github.com/drk1wi/Modlishka/dist/proxy", "-config", *mdlConfig.ProxyDomain+".json")
	cmd.Start()
}
