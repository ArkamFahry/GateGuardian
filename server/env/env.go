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
	Port              string
	EncryptionKey     string
	JwtType           string
	JwtSecret         string
	ClientID          string
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
	port := os.Getenv(constants.Port)
	encryptionKey := os.Getenv(constants.EncryptionKey)
	jwtType := os.Getenv(constants.JwtType)
	jwtSecret := os.Getenv(constants.JwtSecret)
	clientID := os.Getenv(constants.ClientID)

	env := Env{
		DatabaseURL:       dbUrl,
		DatabaseName:      dbName,
		DatabaseNameSpace: dbNameSpace,
		DatabaseUsername:  dbUserName,
		DatabasePassword:  dbPassword,
		Port:              port,
		EncryptionKey:     encryptionKey,
		JwtType:           jwtType,
		JwtSecret:         jwtSecret,
		ClientID:          clientID,
	}

	return env
}
