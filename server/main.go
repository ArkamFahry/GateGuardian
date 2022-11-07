package main

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/envdb"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/routes"
	"github.com/sirupsen/logrus"
)

var VERSION string

type LogUTCFormatter struct {
	logrus.Formatter
}

func (u LogUTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}

func main() {

	constants.VERSION = VERSION

	logrus.SetFormatter(LogUTCFormatter{&logrus.JSONFormatter{}})

	log := logrus.New()
	log.SetFormatter(LogUTCFormatter{&logrus.JSONFormatter{}})

	logLevel := logrus.InfoLevel

	logrus.SetLevel(logLevel)

	// set log level for go-gin middleware
	log.SetLevel(logLevel)

	// initialize envdb provider
	err := envdb.InitEnvDB()
	if err != nil {
		log.Fatal("Error while initializing envdb: ", err)
	}

	// get envs and persists envs to cache
	env.GetEnv()

	// initialize memorydb provider
	err = memorydb.InitMemoryDB()
	if err != nil {
		log.Fatal("Error while initializing memorydb: ", err)
	}

	// initialize maindb provider
	err = maindb.InitMainDB()
	if err != nil {
		log.Fatal("Error while initializing maindb: ", err)
	}

	router := routes.InitRouter(log)
	log.Info("Starting GateGuardian: ", VERSION)
	port, err := env.GetEnvByKey(constants.Port)
	if err != nil {
		log.Error("Error getting port from env: ", err)
		port = "3000"
		log.Info("Switching to default port: ", port)
	}
	log.Info("GateGuardian running at PORT: ", port)

	router.Run(":" + port)
}
