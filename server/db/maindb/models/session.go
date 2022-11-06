package models

type Session struct {
	ID        string `json:"id" gorm:"primaryKey;type:char(36)"`
	UserID    string `json:"user_id"`
	UserAgent string `json:"user_agent"`
	IP        string `json:"ip"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
