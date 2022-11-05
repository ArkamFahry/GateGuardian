package models

type Session struct {
	ID        string `json:"id" sql:"id"`
	UserID    string `json:"user_id" sql:"user_id"`
	UserAgent string `json:"user_agent" sql:"user_agent"`
	IP        string `json:"ip" sql:"ip"`
	CreatedAt int64  `json:"created_at" sql:"created_at"`
	UpdatedAt int64  `json:"updated_at" sql:"updated_at"`
}
