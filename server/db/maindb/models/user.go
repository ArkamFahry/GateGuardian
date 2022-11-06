package models

import (
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/refs"
)

type User struct {
	ID                       string  `json:"id" gorm:"primaryKey;type:char(36)"`
	Email                    string  `json:"email" gorm:"unique"`
	EmailVerifiedAt          *int64  `json:"email_verified_at"`
	Password                 *string `json:"password"`
	SignUpMethods            string  `json:"sign_up_method"`
	UserName                 *string `json:"user_name"`
	FamilyName               *string `json:"family_name"`
	GivenName                *string `json:"given_name"`
	MiddleName               *string `json:"middle_name"`
	NickName                 *string `json:"nick_name"`
	Gender                   *string `json:"gender"`
	BirthDate                *string `json:"birth_date"`
	PhoneNumber              *string `json:"phone_number" gorm:"unique"`
	PhoneNumberVerifiedAt    *int64  `json:"phone_number_verified_at"`
	Picture                  *string `json:"picture"`
	Roles                    string  `json:"roles"`
	RevokedTimestamp         *int64  `json:"revoked_timestamp"`
	IsMultiFactorAuthEnabled *bool   `json:"is_multi_factor_auth_enabled"`
	CreatedAt                int64   `json:"created_at"`
	UpdatedAt                int64   `json:"updated_at"`
	LastLoggedIn             int64   `json:"last_logged_in"`
}

func (user *User) AsAPIUser() *model.User {
	isEmailVerified := user.EmailVerifiedAt != nil
	isPhoneVerified := user.PhoneNumberVerifiedAt != nil

	return &model.User{
		ID:                       user.ID,
		Email:                    user.Email,
		EmailVerified:            isEmailVerified,
		SignupMethods:            user.SignUpMethods,
		GivenName:                user.GivenName,
		FamilyName:               user.FamilyName,
		MiddleName:               user.MiddleName,
		Nickname:                 user.NickName,
		UserName:                 refs.NewStringRef(user.Email),
		Gender:                   user.Gender,
		BirthDate:                user.BirthDate,
		PhoneNumber:              user.PhoneNumber,
		PhoneNumberVerified:      &isPhoneVerified,
		Picture:                  user.Picture,
		Roles:                    strings.Split(user.Roles, ","),
		RevokedTimestamp:         user.RevokedTimestamp,
		IsMultiFactorAuthEnabled: user.IsMultiFactorAuthEnabled,
		CreatedAt:                refs.NewInt64Ref(user.CreatedAt),
		UpdatedAt:                refs.NewInt64Ref(user.UpdatedAt),
	}
}
