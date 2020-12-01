package main

import (
	_ "fmt"
	"os"

	"github.com/duraki/decipiat/cli"
	"github.com/duraki/decipiat/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	cli.InitializeLogging()
	log.Info("Parsing configuration")
	options := cli.ParseConfiguration()

	// Environment?
	environ := os.Getenv("STAGING")
	ssl := false

	if environ != "" {
		ssl = true
	} else {
		ssl = false
	}

	server := web.Server{*options.Host, *options.Port, ssl, *options.CertPath, *options.KeyPath}
	log.Infof("%s %+v\n", "decipiat running via config ...", server.Self())
	server.Run()
}
