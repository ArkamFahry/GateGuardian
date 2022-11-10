package routes

import (
	controllers "github.com/ArkamFahry/GateGuardian/server/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func Auth(router fiber.Router) {
	router.Post("/signup", controllers.Signup)
	router.Post("/login", controllers.Login)
	router.Get("/refresh", controllers.Login)
}