package models

// Env model for db
type Env struct {
	ID        string `json:"id"`
	EnvData   string `json:"env_data"`
	Hash      string `json:"hash"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
}
