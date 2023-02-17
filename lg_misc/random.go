package lg_misc

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomSeed() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(size int) string {
	randomSeed()
	buffer := make([]rune, size)
	for i := range buffer {
		buffer[i] = letters[rand.Intn(len(letters))]
	}
	return string(buffer)
}
