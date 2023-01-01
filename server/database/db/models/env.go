package models

// Database model for env
type Env struct {
	ID        string `json:"_id" gorm:"primaryKey;type:char(36)"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
