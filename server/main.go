package main

import (
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/envdb"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/routes"
	"github.com/sirupsen/logrus"
)

var VERSION string

func main() {

	constants.VERSION = VERSION

	// initialize envdb provider
	err := envdb.InitEnvDB()
	if err != nil {
		logrus.Fatalln("Error while initializing envdb: ", err)
	}

	// get envs and persists envs to cache
	env.GetEnv()

	// initialize memorydb provider
	err = memorydb.InitMemoryDB()
	if err != nil {
		logrus.Fatalln("Error while initializing memorydb: ", err)
	}

	// initialize maindb provider
	err = maindb.InitMainDB()
	if err != nil {
		logrus.Fatalln("Error while initializing maindb: ", err)
	}

	memorydb.Provider.SetSession("basic_auth:uu1239800e03000", "session")

	res, _ := memorydb.Provider.GetSession("basic_auth:uu1239800e03000")

	fmt.Println(res)

	router := routes.InitRouter(logrus.New())
	logrus.Info("Starting GateGuardian: ", VERSION)
	port, err := env.GetEnvByKey(constants.Port)
	if err != nil {
		logrus.Error("Error getting port from env: ", err)
		port = "3000"
		logrus.Info("Switching to default port: ", port)
	}
	logrus.Info("GateGuardian running at PORT: ", port)

	router.Run(":" + port)
}
