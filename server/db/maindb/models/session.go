package models

type Session struct {
	ID        string `json:"id" genji:"id"`
	UserID    string `json:"user_id" genji:"user_id"`
	UserAgent string `json:"user_agent" genji:"user_agent"`
	IP        string `json:"ip" genji:"ip"`
	CreatedAt int64  `json:"created_at" genji:"created_at"`
	UpdatedAt int64  `json:"updated_at" genji:"updated_at"`
}
