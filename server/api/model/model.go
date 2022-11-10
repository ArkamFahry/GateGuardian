package model

type SignupInput struct {
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	GivenName       string `json:"given_name"`
	FamilyName      string `json:"family_name"`
	MiddleName      string `json:"middle_name"`
	NickName        string `json:"nick_name"`
	Gender          string `json:"gender"`
	BirthDate       string `json:"birth_date"`
	Picture         string `json:"picture"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
