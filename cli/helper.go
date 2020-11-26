package cli

import (
	"math/rand"
	"time"
	"os"
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

func CreateDir() string {
	dir := generate()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0644)
	}
	return dir
}
