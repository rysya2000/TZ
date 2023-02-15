package usecase

import (
	"math/rand"
	"time"
)

func Generate(n int) string {
	rand.Seed(time.Now().UnixNano())
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}
