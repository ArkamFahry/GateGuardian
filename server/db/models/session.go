package models

// Session model for db
type Session struct {
	ID           string `json:"id"`
	UserId       string `json:"user_id"`
	UserAgent    string `json:"user_agent"`
	Ip           string `json:"ip"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}
