package models

// Database model for verification_request
type VerificationRequest struct {
	ID          string `json:"_id" gorm:"primaryKey;type:char(36)"`
	Identifier  string `json:"identifier"`
	Email       string `json:"email" gorm:"uniqueIndex:idx_email_identifier;type:varchar(256)"`
	Nonce       string `json:"nonce"`
	RedirectUri string `json:"redirect_uri"`
	ExpiresAt   int64  `json:"expires_at"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
