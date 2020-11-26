package core

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"os/exec"
	"github.com/duraki/decipiat/models"
)

func StartModlishka(mdlConfig models.ModlishkaConfig) {
	file, _ := json.MarshalIndent(mdlConfig," ", " ")
	_ = ioutil.WriteFile(*mdlConfig.ProxyDomain + ".json", file, 0644)

	// modlishkaPath := "~/go/src/github.com/drk1wi/Modlishka/dist/proxy"

	cmd := exec.Command("/Users/daemon1/go/src/github.com/drk1wi/Modlishka/dist/proxy", "-config", *mdlConfig.ProxyDomain + ".json")
	cmd.Start()
}
