package models

// User model for db
type User struct {
	ID                       string  `json:"id"`
	Email                    string  `json:"email"`
	EmailVerifiedAt          *int64  `json:"email_verified_at"`
	Password                 *string `json:"password"`
	SignUpMethods            string  `json:"sign_up_method"`
	GivenName                *string `json:"given_name"`
	FirstName                *string `json:"first_name"`
	MiddleName               *string `json:"middle_name"`
	LastName                 *string `json:"last_name"`
	FullName                 *string `json:"full_name"`
	NickName                 *string `json:"nick_name"`
	Gender                   *string `json:"gender"`
	BirthDate                *string `json:"birth_date"`
	PhoneNumber              *string `json:"phone_number"`
	PhoneNumberVerifiedAt    *int64  `json:"phone_number_verified_at"`
	Picture                  *string `json:"picture"`
	Roles                    string  `json:"roles"`
	DefaultRole              string  `json:"default_role"`
	AssignedRole             string  `json:"assigned_role"`
	AssignedRoles            string  `json:"assigned_roles"`
	RevokedTimestamp         *int64  `json:"revoked_timestamp"`
	IsMultiFactorAuthEnabled *bool   `json:"is_multi_factor_auth_enabled"`
	UpdatedAt                int64   `json:"updated_at"`
	CreatedAt                int64   `json:"created_at"`
	LastLoggedIn             int64   `json:"last_logged_in"`
}
