package models

// WebhookLog model for db
type WebhookLog struct {
	ID         string `json:"id"`
	HttpStatus int64  `json:"http_status"`
	Response   string `json:"response"`
	Request    string `json:"request"`
	WebhookId  string `json:"webhook_id"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}
