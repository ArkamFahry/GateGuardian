package cmd

import (
	"gategaurdian/server/database/memorystore"
	"gategaurdian/server/handlers"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Serve function servers the newly bootstrapped server
func Serve() error {
	var err error

	app := fiber.New()

	app.All("/graphql", handlers.Graphql)
	app.All("/", handlers.PlayGround)

	port := memorystore.RequiredEnvStoreObj.GetRequiredEnv().Port

	err = app.Listen(":" + port)
	if err != nil {
		log.Error("Error starting the server : ", err)
	}

	return err
}
