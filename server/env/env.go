package env

import (
	"os"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env struct {
	DatabaseURL       string
	DatabaseName      string
	DatabaseNameSpace string
	DatabaseUsername  string
	DatabasePassword  string
}

func EnvGet() Env {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Info("Failed To Load .env file switch to os env")
	}

	dbUrl := os.Getenv(constants.DatabaseURL)
	dbName := os.Getenv(constants.DatabaseName)
	dbNameSpace := os.Getenv(constants.DatabaseNameSpace)
	dbUserName := os.Getenv(constants.DatabaseUsername)
	dbPassword := os.Getenv(constants.DatabasePassword)

	env := Env{
		DatabaseURL:       dbUrl,
		DatabaseName:      dbName,
		DatabaseNameSpace: dbNameSpace,
		DatabaseUsername:  dbUserName,
		DatabasePassword:  dbPassword,
	}

	return env
}
