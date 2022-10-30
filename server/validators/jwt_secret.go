package validators

import "errors"

// ValidatePassword to validate the password against the following policy
// min char length: 6
// max char length: 36
// at least one upper case letter
// at least one lower case letter
// at least one digit
// at least one special character
func IsValidJwtSecret(jwtSecret string) error {
	if len(jwtSecret) < 32 {
		return errors.New("jwt secret must be of minimum 32 characters")
	}

	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range jwtSecret {
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

	return errors.New(`jwt secret is not valid. It needs to be at least 32 characters long and needs contain number, uppercase letter, lowercase letter and special character`)
}
