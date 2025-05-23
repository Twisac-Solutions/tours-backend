package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateUsername(name string) string {
	rand.Seed(time.Now().UnixNano())
	base := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	return base + "_" + RandStringRunes(4)
}

func RandStringRunes(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
