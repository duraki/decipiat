package main

import (
    "fmt"
    "github.com/duraki/decipiat/cli"
    "github.com/duraki/decipiat/web"
)

func main() {
    fmt.Println("Decipiat")
    options := cli.ParseConfiguration()
    server := web.Server{*options.Host, *options.Port, false}
    server.PrintConfig()
}
