package models

type ModelList struct {
	User    string
	Session string
	State   string
}

var (
	Model = ModelList{
		User:    "users",
		Session: "sessions",
		State:   "state",
	}
)
