package models

type ModelList struct {
	User    string
	Session string
}

var (
	Model = ModelList{
		User:    "users",
		Session: "sessions",
	}
)
