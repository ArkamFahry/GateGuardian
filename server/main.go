package main

import (
	"github.com/ArkamFahry/GateGuardian/server/api/routes"
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/internal/env"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/sessionstore"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error

	// Environment variable store initialization
	err = envstore.InitEnvStore()
	if err != nil {
		log.Fatal("Failed to initialize env store instance: ", err)
	}

	// Get environment variable and persist to store
	env.GetEnv()

	// Session store initialization
	err = sessionstore.InitSessionStore()
	if err != nil {
		log.Fatal("Failed to initialize session store instance: ", err)
	}

	// Main database initialization and schema migration
	err = db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize main database: ", err)
	}

	app := fiber.New()

	routes.Health(app.Group("/health"))
	routes.Auth(app.Group("/auth"))

	port, err := envstore.Provider.GetEnv(constants.PORT)

	if err != nil {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
