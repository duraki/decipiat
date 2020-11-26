package web

import "fmt"

type Server struct {
    Hostname    string
    Port        int
    Usessl      bool
}

func (srv *Server) PrintConfig() {
    fmt.Printf("%+v\n", srv)
}

func (srv *Server) GetHost() string {
    return srv.Hostname
}
