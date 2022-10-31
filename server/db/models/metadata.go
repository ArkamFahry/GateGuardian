package models

type Metadata struct {
	ID        string         `json:"id"`
	UserId    string         `json:"user_id"`
	Metadata  map[string]any `json:"metadata"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
}
