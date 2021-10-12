package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abdcghijklmnopqrstuvxwyz"

// RandomString ...
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
