package models

// Database model for session
type Session struct {
	ID        string `json:"_id" gorm:"primaryKey;type:char(36)"`
	UserId    string `json:"user_id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
