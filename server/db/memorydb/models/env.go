package models

// Env model for memorydb
type Env struct {
	ID        string `json:"id"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
}
