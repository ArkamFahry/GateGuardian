package models

// Database model for env
type Env struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
