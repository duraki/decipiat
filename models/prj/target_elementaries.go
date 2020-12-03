package models

import (
	"gopkg.in/mgo.v2/bson"
	_ "time"
)

// Project Config struct contains all configuration params per specific project
type (
	TargetElementaries struct {
		ID                  bson.ObjectId `json:"id" bson:"_id,omitempty"`
		ProjectID           bson.ObjectId `json:"prjid" bson:"_id,omitempty"`
		CpvUuid             string        `json:"cpvUuid" bson:"cpvUuid"`
		TargetURI           string        `json:"targetUri" bson:"targetUri"`
		TargetURIVerified   bool          `json:"tuv" bson:"tuv"` // can attack this uri?
		RedirectURI         string        `json:"redirectUri" bson:"redirectUri"`
		RedirectURIVerified bool          `json:"ruv" bson:"ruv"` // can point to this uri?
	}
)

const (
	CollectionTargetElementaries = "prj.elementaries"
)
