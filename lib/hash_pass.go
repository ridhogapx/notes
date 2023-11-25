package lib

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {
		fmt.Println("Failed to hashing password:", err)
		return ""
	}

	return string(hash)
}
