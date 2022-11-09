package controllers

import "github.com/gofiber/fiber/v2"

func Health(c *fiber.Ctx) error {
	return c.SendString("Api is healthy")
}
