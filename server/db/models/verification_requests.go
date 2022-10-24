package models

// VerificationRequest model for db
type VerificationRequest struct {
	ID          string `json:"id"`
	Token       string `json:"token"`
	Identifier  string `json:"identifier"`
	ExpiresAt   int64  `json:"expires_at"`
	Email       string `json:"email"`
	Nonce       string `json:"nonce"`
	RedirectUri string `json:"redirect_uri"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
