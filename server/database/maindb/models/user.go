package models

// Database model for user
type User struct {
	ID                        string  `json:"id"`
	Email                     string  `json:"email"`
	EmailVerified             *bool   `json:"email_verified"`
	EmailVerifiedAt           *int64  `json:"email_verified_at"`
	PasswordHash              string  `json:"password_hash"`
	Roles                     string  `json:"roles"`
	PhoneNumber               *string `json:"phone_number"`
	PhoneNumberVerified       *bool   `json:"phone_number_verified"`
	PhoneNumberVerifiedAt     *int64  `json:"phone_number_verified_at"`
	Name                      *string `json:"name"`
	GivenName                 *string `json:"given_name"`
	MiddleName                *string `json:"middle_name"`
	FamilyName                *string `json:"family_name"`
	BirthDate                 *string `json:"birth_date"`
	Gender                    *string `json:"gender"`
	Avatar                    *string `json:"avatar"`
	MultiFactorAuthEnabled    *bool   `json:"multi_factor_auth_enabled"`
	MultiFactorAuthEnabledAt  *int64  `json:"multi_factor_auth_enabled_at"`
	MultiFactorAuthDisabledAt *int64  `json:"multi_factor_auth_disabled_at"`
	AccessRevoked             *bool   `json:"access_revoked"`
	AccessRevokedAt           *int64  `json:"access_revoked_at"`
	AccessUnRevokedAt         *int64  `json:"access_un_revoked_at"`
	Archived                  *bool   `json:"archived"`
	ArchivedAt                *int64  `json:"archived_at"`
	UnArchivedAt              *int64  `json:"un_archived_at"`
	MetaData                  *string `json:"meta_data"`
	CreatedAt                 int64   `json:"created_at"`
	UpdatedAt                 int64   `json:"updated_at"`
}
