package model

type SignupInput struct {
	Email           string  `json:"email" validate:"required"`
	Password        string  `json:"password" validate:"required"`
	ConfirmPassword string  `json:"confirm_password" validate:"required"`
	GivenName       *string `json:"given_name"`
	FamilyName      *string `json:"family_name"`
	MiddleName      *string `json:"middle_name"`
	NickName        *string `json:"nick_name"`
	Gender          *string `json:"gender"`
	BirthDate       *string `json:"birth_date"`
	Picture         *string `json:"picture"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	Id         string  `json:"id"`
	Email      string  `json:"email"`
	GivenName  *string `json:"given_name"`
	FamilyName *string `json:"family_name"`
	MiddleName *string `json:"middle_name"`
	NickName   *string `json:"nick_name"`
	Gender     *string `json:"gender"`
	BirthDate  *string `json:"birth_date"`
	Picture    *string `json:"picture"`
	CreatedAt  int64   `json:"created_at"`
	UpdatedAt  int64   `json:"updated_at"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

type AuthResponse struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpires  int64  `json:"access_token_expires"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpires int64  `json:"refresh_token_expires"`
	IdToken             string `json:"id_token"`
	IdTokenExpires      int64  `json:"id_token_expires"`
}
