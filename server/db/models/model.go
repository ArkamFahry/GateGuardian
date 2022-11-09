package models

type ModelList struct {
	User string
}

var (
	Models = ModelList{
		User: "users",
	}
)
