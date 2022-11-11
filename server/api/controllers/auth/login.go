package controllers

import (
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/api/model"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/db/models"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/sessionstore"
	"github.com/ArkamFahry/GateGuardian/server/tokens"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
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

	session := models.Session{
		UserId:    user.Id,
		UserAgent: string(ctx.UserAgent()),
		Ip:        c.IP(),
	}

	err = db.Provider.AddSession(ctx, session)
	if err != nil {
		log.Debug("error inserting session to db : ", err)
	}

	tokens, err := tokens.CreateAuthTokens(user, c.Hostname())
	if err != nil {
		log.Debug("error creating auth tokens : ", err)
	}

	rt_token_hash, err := crypto.EncryptData(tokens.RefreshToken)
	if err != nil {
		log.Debug("error hashing refresh token : ", err)
	}
	sessionstore.Provider.SetSession(user.Id, rt_token_hash)

	return c.Status(201).JSON(fiber.Map{"message": "successfully login", "tokens": tokens})
}
