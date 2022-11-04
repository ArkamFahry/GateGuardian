package models

type Metadata struct {
	ID        string         `json:"id" genji:"id"`
	UserId    string         `json:"user_id" genji:"user_id"`
	Metadata  map[string]any `json:"metadata" genji:"metadata"`
	CreatedAt int64          `json:"created_at" genji:"created_at"`
	UpdatedAt int64          `json:"updated_at" genji:"updated_at"`
}
