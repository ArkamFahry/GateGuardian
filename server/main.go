package main

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/routes"
	"github.com/sirupsen/logrus"
)

var VERSION string

func main() {

	constants.VERSION = VERSION

	// initialize memorydb provider
	err := memorydb.InitMemoryDB()
	if err != nil {
		logrus.Fatalln("Error while initializing memorydb: ", err)
	}

	// persists all envs to cache
	env.PersistEnvToCache()

	// initialize maindb provider
	err = maindb.InitMainDB()
	if err != nil {
		logrus.Fatalln("Error while initializing maindb: ", err)
	}

	router := routes.InitRouter(logrus.New())
	logrus.Info("Starting GateGuardian: ", VERSION)
	port, err := memorydb.Provider.GetEnvByKey(constants.Port)
	if err != nil {
		logrus.Error("Error getting port from env: ", err)
		port = "3000"
		logrus.Info("Switching to default port: ", port)
	}
	logrus.Info("GateGuardian running at PORT: ", port)

	router.Run(":" + port)
}
