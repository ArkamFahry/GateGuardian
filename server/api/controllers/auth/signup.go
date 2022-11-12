package controllers

import (
	"fmt"
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/api/model"
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/ArkamFahry/GateGuardian/server/validators"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func ValidateSignupInput(signUpInput model.SignupInput) []*model.ErrorResponse {
	var validate = validator.New()
	var errors []*model.ErrorResponse
	err := validate.Struct(signUpInput)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)
		}
	}

	return errors
}

func Signup(c *fiber.Ctx) error {
	var params model.SignupInput
	var user models.User

	c.BodyParser(&params)

	ctx := c.Context()

	errors := ValidateSignupInput(params)
	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	params.Email = strings.ToLower(params.Email)

	existingUser, err := db.Provider.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		log.Debug("Failed to get user by email: ", err)
	}

	if existingUser.Id != "" {
		return c.Status(400).JSON(fiber.Map{
			"error":  "user with the email already exists",
			"reason": fmt.Sprintf("%s has already signed up", existingUser.Email)})
	} else {
		user.Email = params.Email
	}

	if params.Password != params.ConfirmPassword {
		return c.Status(400).JSON(fiber.Map{"error": "Password and Confirm Password are not equal"})
	} else {
		password, _ := crypto.EncryptData(params.Password)
		user.Password = &password
	}

	if len(params.Roles) > 0 {
		AllowedRoles, err := envstore.Provider.GetEnv(constants.AllowedRoles)
		allowedRoles := strings.Split(AllowedRoles, ",")
		if err != nil {
			log.Debug("can't fetch allowed roles from env")
		}
		if validators.IsValidRoles(allowedRoles, params.Roles) {
			user.Roles = strings.Join(params.Roles, ",")
		} else {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid user role", "reason": "user role doesn't exist on allowed role list"})
		}
	} else {
		roles, err := envstore.Provider.GetEnv(constants.Roles)
		if err != nil {
			log.Debug("can't fetch roles from env")
		}
		user.Roles = roles
	}

	if params.GivenName != nil {
		user.GivenName = params.GivenName
	}

	if params.FamilyName != nil {
		user.FamilyName = params.FamilyName
	}

	if params.MiddleName != nil {
		user.MiddleName = params.MiddleName
	}

	if params.NickName != nil {
		user.NickName = params.NickName
	}

	if params.Gender != nil {
		user.Gender = params.Gender
	}

	if params.BirthDate != nil {
		user.BirthDate = params.BirthDate
	}

	if params.Picture != nil {
		user.Picture = params.Picture
	}

	user, err = db.Provider.AddUser(ctx, user)
	if err != nil {
		log.Debug("Failed to insert user to db: ", err)
	}

	return c.Status(201).JSON(fiber.Map{"message": "successfully signed up", "user": user.AsAPIUser()})
}
