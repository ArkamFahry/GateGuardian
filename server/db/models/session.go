package models

type Session struct {
	Id        string `gorm:"primaryKey;type:char(36)" json:"id" db:"id"`
	UserId    string `json:"user_id" db:"user_id"`
	UserAgent string `json:"user_agent" db:"user_agent"`
	Ip        string `json:"ip" db:"ip"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}
