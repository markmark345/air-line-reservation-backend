package utils

import (
	"math/rand"
	"strings"
	"time"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ")

func RandomInt(min, max int64) int64 {
	return randomGenerator.Int63n(max - min + 1) + min 
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(charset)

	for i := 0; i < n; i++ {
		c := charset[randomGenerator.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomGender() string {
	gender := []string{"M", "F", "N"}
	n := len(gender)
	return gender[randomGenerator.Intn(n)]
}

func RandomEmail() string {
	username := RandomString(8)
	domain := "gmail.com"
	return username + "@" + domain
}
