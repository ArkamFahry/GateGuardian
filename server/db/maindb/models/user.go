package models

type User struct {
	ID                       string  `json:"id" sql:"id"`
	Email                    string  `json:"email" sql:"email"`
	EmailVerifiedAt          *int64  `json:"email_verified_at" sql:"email_verified_at"`
	Password                 *string `json:"password" sql:"password"`
	SignUpMethods            string  `json:"sign_up_method" sql:"sign_up_method"`
	UserName                 *string `json:"user_name" sql:"user_name"`
	FamilyName               *string `json:"family_name" sql:"family_name"`
	GivenName                *string `json:"given_name" sql:"given_name"`
	MiddleName               *string `json:"middle_name" sql:"middle_name"`
	NickName                 *string `json:"nick_name" sql:"nick_name"`
	Gender                   *string `json:"gender" sql:"gender"`
	BirthDate                *string `json:"birth_date" sql:"birth_date"`
	PhoneNumber              *string `json:"phone_number" sql:"phone_number"`
	PhoneNumberVerifiedAt    *int64  `json:"phone_number_verified_at" sql:"phone_number_verified_at"`
	Picture                  *string `json:"picture" sql:"picture"`
	Roles                    string  `json:"roles" sql:"roles"`
	RevokedTimestamp         *int64  `json:"revoked_timestamp" sql:"revoked_timestamp"`
	IsMultiFactorAuthEnabled *bool   `json:"is_multi_factor_auth_enabled" sql:"is_multi_factor_auth_enabled"`
	CreatedAt                int64   `json:"created_at" sql:"created_at"`
	UpdatedAt                int64   `json:"updated_at" sql:"updated_at"`
	LastLoggedIn             int64   `json:"last_logged_in" sql:"last_logged_in"`
}
