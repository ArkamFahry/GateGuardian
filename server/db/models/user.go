package models

type User struct {
	Id         string `db:"id"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	GivenName  string `db:"given_name"`
	FamilyName string `db:"family_name"`
	MiddleName string `db:"middle_name"`
	NickName   string `db:"nick_name"`
	Gender     string `db:"gender"`
	CreatedAt  int64  `db:"created_at"`
	UpdatedAt  int64  `db:"updated_at"`
}
