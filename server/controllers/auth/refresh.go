package controllers

import "github.com/gofiber/fiber/v2"

func Refresh(c *fiber.Ctx) error {
	return c.SendString("Refresh is healthy")
}
