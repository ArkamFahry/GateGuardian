package cmd

import (
	"github.com/gofiber/fiber/v2"
)

func Serve() error {
	var err error

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err = app.Listen(":3000")

	return err
}
