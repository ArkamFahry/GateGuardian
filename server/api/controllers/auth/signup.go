package controllers

import (
	"github.com/ArkamFahry/GateGuardian/server/api/model"
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	var signup model.SignupInput
	var user models.User
	c.BodyParser(&signup)

	userExist, _ := db.Provider.GetUserByEmail(c.Context(), signup.Email)

	if userExist.Email == signup.Email {
		return c.Status(400).JSON("user with the email already exists")
	} else {
		user.Email = signup.Email
	}

	if signup.Password != signup.ConfirmPassword {
		return c.Status(400).JSON("Password and Confirm Password are not equal")
	} else {
		user.Password = signup.Password
	}

	db.Provider.AddUser(c.Context(), user)

	return nil
}
