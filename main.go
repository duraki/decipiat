package main

import (
	_ "fmt"
	"github.com/duraki/decipiat/cli"
	"github.com/duraki/decipiat/web"
	"log"
)

func main() {
	options := cli.ParseConfiguration()
	cli.Usage()

	server := web.Server{*options.Host, *options.Port, false}
	log.Printf("%s %+v\n", "decipiat running via config ...", server.Self())
	server.Run()
}
