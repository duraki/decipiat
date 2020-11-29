package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Project struct contains data about specific project
type (
	Project struct {
		ID         bson.ObjectId `json:"id" bson:"_id,omitempty"` // user id
		Name       string        `json:"name" bson:"name"`        // email
		InternalId string        `json:"int" bson:"int"`          // password (SHA256 || Bcrypt)
		CpvUuid    string        `json:"cpvUuid" bson:"cpvUuid"`  // type => admin, org, user
		Status     int           `json:"status" bson:"s"`         // 0 = inactive, 1 = active, -1 = deleted, 66 = banned
		UserID     bson.ObjectId `json:"userId" bson:"_userid,omitempty"`
		CreatedAt  time.Time     `json:"createdAt,omitempty" bson:"createdAt"` // ...
		UpdatedAt  time.Time     `json:"updatedAt,omitempty" bson:"updatedAt"` // ...
	}
)

const (
	CollectionProject = "prj"
)
