package main

import (
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/sessionstore"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error

	err = envstore.InitEnvStore()
	if err != nil {
		log.Fatal("Failed to initialize env store instance: ", err)
	}

	env.GetEnv()

	err = sessionstore.InitSessionStore()
	if err != nil {
		log.Fatal("Failed to initialize session store instance: ", err)
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize main database: ", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":" + "3100"))
}
