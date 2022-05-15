package util

import (
	"golang.org/x/crypto/bcrypt"
)

// returns the Bcrypt hash of a password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// checks if hashes are the same
func CheckPassword(hashToVerify string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashToVerify), []byte(password))
}
