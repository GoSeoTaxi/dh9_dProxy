package model

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}
	return string(result)
}
