package validators

import "net/mail"

// validates email
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
