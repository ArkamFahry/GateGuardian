package main

import (
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/sessionstore"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error

	err = envstore.InitEnvStore()
	if err != nil {
		log.Debug("Failed to initialize env store instance: ", err)
	}

	err = sessionstore.InitSessionStore()
	if err != nil {
		log.Debug("Failed to initialize session store instance: ", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
