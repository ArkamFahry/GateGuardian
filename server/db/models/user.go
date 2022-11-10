package models

type User struct {
	Id         string `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	GivenName  string `json:"given_name" db:"given_name"`
	FamilyName string `json:"family_name" db:"family_name"`
	MiddleName string `json:"middle_name" db:"middle_name"`
	NickName   string `json:"nick_name" db:"nick_name"`
	Gender     string `json:"gender" db:"gender"`
	BirthDate  string `json:"birth_date" db:"birth_date"`
	Picture    string `json:"picture" db:"picture"`
	CreatedAt  int64  `json:"created_at" db:"created_at"`
	UpdatedAt  int64  `json:"updated_at" db:"updated_at"`
}
