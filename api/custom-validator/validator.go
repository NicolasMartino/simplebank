package customvalidator

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func PasswordValidator(password string, minEntropy float64) bool {
	entropy := passwordvalidator.GetEntropy(password)
	return entropy > minEntropy
}
