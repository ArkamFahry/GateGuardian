package models

// Env model for db
type Env struct {
	ID        string `json:"id" gorm:"id"`
	Data      string `json:"data" gorm:"data"`
	Hash      string `json:"hash" gorm:"hash"`
	UpdatedAt int64  `json:"updated_at" gorm:"created_at"`
	CreatedAt int64  `json:"created_at" gorm:"updated_at"`
}
