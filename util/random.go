package util

import (
	"math/rand"
	"time"
)

var key = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomString(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += string(key[rand.Intn(len(key))])
	}
	return res
}

func RandomPassword(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += string(key[rand.Intn(len(key))])
	}
	return res
}

func RandomEmail(n int) string {
	res := ""

	for i := 0; i < n; i++ {
		res += string(key[rand.Intn(len(key))])
	}
	return res + "@mail.com"
}

func RandomTime(hours int) time.Time {
	hour := rand.Intn(hours)
	return time.Now().Add(time.Duration(hour) * time.Hour)
}
