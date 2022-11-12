package models

import (
	"encoding/json"
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/api/model"
)

type User struct {
	Id         string  `gorm:"primaryKey;type:char(36)" json:"id" db:"id"`
	Email      string  `gorm:"unique" json:"email" db:"email"`
	Password   *string `json:"password" db:"password"`
	GivenName  *string `json:"given_name" db:"given_name"`
	FamilyName *string `json:"family_name" db:"family_name"`
	MiddleName *string `json:"middle_name" db:"middle_name"`
	NickName   *string `json:"nick_name" db:"nick_name"`
	Gender     *string `json:"gender" db:"gender"`
	BirthDate  *string `json:"birth_date" db:"birth_date"`
	Picture    *string `json:"picture" db:"picture"`
	Roles      string  `json:"roles" db:"roles"`
	CreatedAt  int64   `json:"created_at" db:"created_at"`
	UpdatedAt  int64   `json:"updated_at" db:"updated_at"`
}

func (user *User) AsAPIUser() *model.User {
	return &model.User{
		Id:         user.Id,
		Email:      user.Email,
		Roles:      strings.Split(user.Roles, ","),
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
