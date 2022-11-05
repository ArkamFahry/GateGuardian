package models

// Env model for memorydb
type Env struct {
	ID        string `json:"id" genji:"id"`
	Data      string `json:"data" genji:"data"`
	CreatedAt int64  `json:"created_at" genji:"created_at"`
	UpdatedAt int64  `json:"updated_at" genji:"updated_at"`
}
