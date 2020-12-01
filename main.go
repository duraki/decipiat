package main

import (
	_ "fmt"

	"github.com/duraki/decipiat/cli"
	"github.com/duraki/decipiat/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	cli.InitializeLogging()
	log.Info("Parsing configuration")
	options := cli.ParseConfiguration()

	server := web.Server{*options.Host, *options.Port, *options.Ssl, *options.CertPath, *options.KeyPath}
	log.Infof("%s %+v\n", "decipiat running via config ...", server.Self())
	server.Run()
}
