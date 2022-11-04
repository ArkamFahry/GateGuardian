package validators

import (
	"errors"
)

func IsValidPassword(password string) error {
	if len(password) < 8 || len(password) > 36 {
		return errors.New("password must be of minimum 8 characters and maximum 36 characters")
	}

	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpperCase = true
		} else if char >= 'a' && char <= 'z' {
			hasLowerCase = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		} else {
			hasSpecialChar = true
		}
	}

	isValid := hasUpperCase && hasLowerCase && hasDigit && hasSpecialChar

	if isValid {
		return nil
	}

	return errors.New(`password is not valid. It needs to be at least 8 characters long and contain at least one number, one uppercase letter, one lowercase letter and one special character`)
}
