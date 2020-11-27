package models

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"time"
)

const (
	CollectionUser = "usersCollection"
)

type (
	Subscription struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty`     // user id
		Email     string        `json:"email" bson:"email"`         // email
		Password  string        `json:"password" bson:"password`    // password (SHA256 || Bcrypt)
		Type      string        `json:"type" bson:"usertype"`       // type => admin, org, user
		CreatedAt time.Time     `json:"createdAt" bson:"createdAt"` // ...
		UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"` // ...
		Status    int           `json:"status" bson:"s"`            // 0 = inactive, 1 = active, -1 = deleted, 66 = banned
	}
)
