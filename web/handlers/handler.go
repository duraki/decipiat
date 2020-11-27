package handlers

import (
	"gopkg.in/mgo.v2"
)

/**
 * Root-level handler for DBMS, * Controller extension, * & alike.
 */

var GlobalConfig = Globals{}
var DatabaseName string = "decipiat"

type (
	Globals struct {
		DB *mgo.Session
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
