package cmd

import (
	"gategaurdian/server/handlers"

	"github.com/gofiber/fiber/v2"
)

// Serve function servers the newly bootstrapped server
func Serve() error {
	var err error

	app := fiber.New()

	app.All("/graphql", handlers.Graphql)
	app.All("/", handlers.PlayGround)

	err = app.Listen(":3000")

	return err
}
