package models

type Permissions struct {
	ID        string `json:"id"`
	UserId    string `json:"user_id"`
	NameSpace string `json:"name_space"`
	Object    string `json:"object"`
	Relation  string `json:"relation"`
	Subject   string `json:"subject"`
}
