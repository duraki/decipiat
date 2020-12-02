package core

import (
	"math/rand"
	"os"
	"time"

	"github.com/duraki/decipiat/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generate() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 15)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func createDir(projectName string) string {
	dir := projectName + generate()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0644)
	}
	return dir
}

// InitializeProject function handles project initialization. Later on we will add connection to the database and such
func InitializeProject(project models.Project) {
	// Create Directory for it for now, later will probably be some database
	createDir(project.Name)
}
