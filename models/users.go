package models

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"time"
)

// Project struct contains data about specific project
type (
	User struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty`     // user id
		Email     string        `json:"email" bson:"email"`         // email
		Password  string        `json:"password" bson:"password`    // password (SHA256 || Bcrypt)
		Type      string        `json:"type" bson:"usertype"`       // type => admin, org, user
		CreatedAt time.Time     `json:"createdAt" bson:"createdAt"` // ...
		UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"` // ...
		Status    int           `json:"status" bson:"s"`            // 1 = active, 0 = deleted, 66 = banned
	}
)

func HashPassword(password string) (string, error) {
	log.Infof("Hashing password total bytes len %s", len(password))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func HashCompare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
