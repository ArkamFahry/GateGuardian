package models

type User struct {
	ID                       string  `json:"id" genji:"id"`
	Email                    string  `json:"email" genji:"email"`
	EmailVerifiedAt          *int64  `json:"email_verified_at" genji:"email_verified_at"`
	Password                 *string `json:"password" genji:"password"`
	SignUpMethods            string  `json:"sign_up_method" genji:"sign_up_method"`
	UserName                 *string `json:"user_name" genji:"user_name"`
	FamilyName               *string `json:"family_name" genji:"family_name"`
	GivenName                *string `json:"given_name" genji:"given_name"`
	MiddleName               *string `json:"middle_name" genji:"middle_name"`
	NickName                 *string `json:"nick_name" genji:"nick_name"`
	Gender                   *string `json:"gender" genji:"gender"`
	BirthDate                *string `json:"birth_date" genji:"birth_date"`
	PhoneNumber              *string `json:"phone_number" genji:"phone_number"`
	PhoneNumberVerifiedAt    *int64  `json:"phone_number_verified_at" genji:"phone_number_verified_at"`
	Picture                  *string `json:"picture" genji:"picture"`
	Roles                    string  `json:"roles" genji:"roles"`
	RevokedTimestamp         *int64  `json:"revoked_timestamp" genji:"revoked_timestamp"`
	IsMultiFactorAuthEnabled *bool   `json:"is_multi_factor_auth_enabled" genji:"is_multi_factor_auth_enabled"`
	CreatedAt                int64   `json:"created_at" genji:"created_at"`
	UpdatedAt                int64   `json:"updated_at" genji:"updated_at"`
	LastLoggedIn             int64   `json:"last_logged_in" genji:"last_logged_in"`
}
