package main

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode(length int) string{
	code := make([]byte, length)

	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)

}