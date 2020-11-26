package web

import (
	"fmt"
	_ "github.com/duraki/decipiat/web/handlers"
)

type Server struct {
	Hostname string
	Port     int
	Usessl   bool
}

var isRunning bool = false /* is current serveload up and running */

func (srv *Server) PrintConfig() {
	fmt.Printf("%+v\n", srv)
}

func (srv *Server) Self() *Server {
	return srv
}

func (srv *Server) GetHost() string {
	return srv.Hostname
}

func (srv *Server) GetPort() int {
	return srv.Port
}

func (srv *Server) GetUseSsl() bool {
	return srv.Usessl
}

func (srv *Server) GetIsRunning() bool {
	return isRunning
}

func (srv *Server) Run() {
	router := Init()
	router.Logger.Fatal(router.Start(":8000")) // todo: fix to use srv*
	isRunning = true
}
