package web

import (
	"fmt"
	"strconv"

	_ "github.com/duraki/decipiat/web/handlers"
)

type Server struct {
	Hostname string
	Port     int
	Usessl   bool
	CertPath string
	KeyPath  string
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
	isRunning = true
	listeningAddress := srv.Hostname + ":" + strconv.Itoa(srv.Port)

	if !srv.Usessl {
		router.Logger.Fatal(router.Start(listeningAddress))
	} else {
		router.Logger.Fatal(router.StartTLS(listeningAddress, srv.CertPath, srv.KeyPath))
	}
}
