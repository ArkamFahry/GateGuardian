package controllers

import (
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/api/model"
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func ValidateLoginInput(loginInput model.LoginInput) []*model.ErrorResponse {
	var validate = validator.New()
	var errors []*model.ErrorResponse
	err := validate.Struct(loginInput)
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

func Login(c *fiber.Ctx) error {
	var params model.LoginInput

	c.BodyParser(&params)

	ctx := c.Context()

	errors := ValidateLoginInput(params)
	if errors != nil {
		return c.Status(400).JSON(errors)
	}

	params.Email = strings.ToLower(params.Email)

	user, err := db.Provider.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad user credentials", "reason": "wrong user email"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(params.Password))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad user credentials", "reason": "wrong user password"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "successfully login", "user": user.AsAPIUser()})
}
