package web

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

type Server struct {
    Hostname    string
    Port        int
    Usessl      bool
    isRunning	bool 	/* is current serveload up and running */
}

func (srv *Server) PrintConfig() {
    fmt.Printf("%+v\n", srv)
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
	return srv.isRunning
}

func (srv *Server) Serve() {
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}