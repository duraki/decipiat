package main

import (
    "fmt"
    "github.com/duraki/decipiat/web"
)

func main() {
    fmt.Println("Decipiat")
    server := web.Server{"127.0.0.1", 8080, false}
    fmt.Println(server.GetHost())
}
