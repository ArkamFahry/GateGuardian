package routes

import (
	controllers "github.com/ArkamFahry/GateGuardian/server/controllers/health"
	"github.com/gofiber/fiber/v2"
)

func Health(router fiber.Router) {
	router.Get("/", controllers.Health)
}
