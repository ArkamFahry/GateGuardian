package main

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/routes"
	"github.com/sirupsen/logrus"
)

var VERSION string

func main() {

	constants.VERSION = VERSION

	// initialize db provider
	err := db.InitDB()
	if err != nil {
		logrus.Fatalln("Error while initializing db: ", err)
	}

	router := routes.InitRouter(logrus.New())
	logrus.Info("Starting GateGuardian: ", VERSION)
	port := env.EnvGet().Port
	logrus.Info("GateGuardian running at PORT: ", port)

	router.Run(":" + port)
}
