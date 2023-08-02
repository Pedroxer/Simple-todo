package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed hashing password", err)
	}
	return string(hashedPassword)
}
