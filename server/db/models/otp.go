package models

// OTP model for database
type OTP struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Otp       string `json:"otp"`
	ExpiresAt int64  `json:"expires_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
