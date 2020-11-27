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
	cli.Usage()

	server := web.Server{*options.Host, *options.Port, false}
	log.Infof("%s %+v\n", "decipiat running via config ...", server.Self())
	server.Run()
}
