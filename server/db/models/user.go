package models

import (
	"encoding/json"

	"github.com/ArkamFahry/GateGuardian/server/api/model"
)

type User struct {
	Id           string  `json:"id" db:"id"`
	Email        string  `json:"email" db:"email"`
	Password     *string `json:"password" db:"password"`
	GivenName    *string `json:"given_name" db:"given_name"`
	FamilyName   *string `json:"family_name" db:"family_name"`
	MiddleName   *string `json:"middle_name" db:"middle_name"`
	NickName     *string `json:"nick_name" db:"nick_name"`
	Gender       *string `json:"gender" db:"gender"`
	BirthDate    *string `json:"birth_date" db:"birth_date"`
	Picture      *string `json:"picture" db:"picture"`
	AllowedRoles string  `json:"allowed_roles" db:"allowed_roles"`
	DefaultRoles string  `json:"default_roles" db:"default_roles"`
	DefaultRole  string  `json:"default_role" db:"default_role"`
	CreatedAt    int64   `json:"created_at" db:"created_at"`
	UpdatedAt    int64   `json:"updated_at" db:"updated_at"`
}

func (user *User) AsAPIUser() *model.User {
	return &model.User{
		Id:         user.Id,
		Email:      user.Email,
		GivenName:  user.GivenName,
		FamilyName: user.FamilyName,
		MiddleName: user.MiddleName,
		NickName:   user.NickName,
		Gender:     user.Gender,
		BirthDate:  user.BirthDate,
		Picture:    user.Picture,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func (user *User) ToMap() map[string]interface{} {
	res := map[string]interface{}{}
	data, _ := json.Marshal(user) // Convert to a json string
	json.Unmarshal(data, &res)    // Convert to a map
	return res
}
