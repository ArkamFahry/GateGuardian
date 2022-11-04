package models

// Env model for db
type Env struct {
	ID        string `json:"id" genji:"id"`
	Data      string `json:"data" genji:"data"`
	Hash      string `json:"hash" genji:"hash"`
	UpdatedAt int64  `json:"updated_at" genji:"created_at"`
	CreatedAt int64  `json:"created_at" genji:"updated_at"`
}
