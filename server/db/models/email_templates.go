package models

// EmailTemplate model for database
type EmailTemplate struct {
	ID        string `json:"id"`
	EventName string `json:"event_name"`
	Subject   string `json:"subject"`
	Template  string `json:"template"`
	Design    string `json:"design"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
