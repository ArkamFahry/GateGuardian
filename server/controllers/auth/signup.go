package controllers

import (
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	var signupUser models.SignUpUser
	c.BodyParser(&signupUser)

	user := models.User{
		Email:      signupUser.Email,
		Password:   signupUser.Password,
		GivenName:  signupUser.GivenName,
		FamilyName: signupUser.FamilyName,
		MiddleName: signupUser.MiddleName,
		NickName:   signupUser.NickName,
		Gender:     signupUser.Gender,
		BirthDate:  signupUser.BirthDate,
		Picture:    signupUser.Picture,
	}

	db.Provider.AddUser(c.Context(), user)

	return nil
}
