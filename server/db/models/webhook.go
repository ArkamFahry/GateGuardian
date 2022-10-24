package models

// Webhook model for db
type Webhook struct {
	ID        string `json:"id"`
	EventName string `json:"event_name"`
	EndPoint  string `json:"endpoint"`
	Headers   string `json:"headers"`
	Enabled   bool   `json:"enabled"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
