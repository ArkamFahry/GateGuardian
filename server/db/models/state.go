package models

type State struct {
	Id            string `json:"id" db:"id"`
	UserId        string `json:"user_id" db:"user_id"`
	CodeChallenge string `json:"code_challenge" db:"code_challenge"`
	AuthCode      string `json:"auth_code" db:"auth_code"`
	CreatedAt     int64  `json:"created_at" db:"created_at"`
	UpdatedAt     int64  `json:"updated_at" db:"updated_at"`
}
