package models

// Session model for db
type Session struct {
	ID           string `json:"id" genji:"id"`
	UserId       string `json:"user_id" genji:"user_id"`
	UserAgent    string `json:"user_agent" genji:"user_agent"`
	Ip           string `json:"ip" genji:"ip"`
	RefreshToken string `json:"refresh_token" genji:"refresh_token"`
	CreatedAt    int64  `json:"created_at" genji:"created_at"`
	UpdatedAt    int64  `json:"updated_at" genji:"updated_at"`
}
