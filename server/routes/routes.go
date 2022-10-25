package routes

import (
	"github.com/ArkamFahry/GateGuardian/server/handlers"
	"github.com/ArkamFahry/GateGuardian/server/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRouter(log *logrus.Logger) *gin.Engine {
	router := gin.New()

	router.Use(middlewares.Logger(log), gin.Recovery())

	router.POST("/graphql", handlers.GraphqlHandler())
	router.GET("/playground", handlers.PlaygroundHandler())

	return router
}
