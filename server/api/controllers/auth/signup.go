package controllers

import (
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/api/model"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Signup(c *fiber.Ctx) error {
	var params model.SignupInput
	var user models.User
	c.BodyParser(&params)

	params.Email = strings.ToLower(params.Email)

	existingUser, err := db.Provider.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		log.Debug("Failed to get user by email: ", err)
	}

	if existingUser.Id != "" {
		return c.Status(400).JSON("user with the email already exists")
	} else {
		user.Email = params.Email
	}

	if params.Password != params.ConfirmPassword {
		return c.Status(400).JSON("Password and Confirm Password are not equal")
	} else {
		password, _ := crypto.EncryptPassword(params.Password)
		user.Password = password
	}

	if params.GivenName != "" {
		user.GivenName = params.GivenName
	}

	if params.FamilyName != "" {
		user.FamilyName = params.FamilyName
	}

	if params.MiddleName != "" {
		user.MiddleName = params.MiddleName
	}

	if params.NickName != "" {
		user.NickName = params.NickName
	}

	if params.Gender != "" {
		user.Gender = params.Gender
	}

	if params.BirthDate != "" {
		user.BirthDate = params.BirthDate
	}

	if params.Picture != "" {
		user.Picture = params.Picture
	}

	db.Provider.AddUser(c.Context(), user)

	return nil
}
