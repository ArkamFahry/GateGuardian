package env

import (
	"os"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/google/uuid"
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
	JwtPrivateKey     string
	JwtPublicKey      string
	ClientID          string
}

func EnvGet() Env {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Info("Failed to load .env file: ", err, "|switched to os env|")
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
	jwtPrivateKey := os.Getenv(constants.JwtPrivateKey)
	jwtPublicKey := os.Getenv(constants.JwtPublicKey)
	clientID := os.Getenv(constants.ClientID)

	if clientID == "" {
		clientID = uuid.New().String()
	}

	if jwtType == "" {
		jwtType = "RS256"
	}

	if jwtPrivateKey == "" || jwtPublicKey == "" {
		if crypto.IsRSA(jwtType) {
			_, jwtPrivateKey, jwtPublicKey, _, _ = crypto.NewRSAKey(jwtType, clientID)
		} else if crypto.IsECDSA(jwtType) {
			_, jwtPrivateKey, jwtPublicKey, _, _ = crypto.NewECDSAKey(jwtType, clientID)
		}
	}

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
		JwtPrivateKey:     jwtPrivateKey,
		JwtPublicKey:      jwtPublicKey,
		ClientID:          clientID,
	}

	memorydb.Provider.AddEnv(constants.DatabaseURL, env.DatabaseURL)
	memorydb.Provider.AddEnv(constants.DatabaseName, env.DatabaseName)
	memorydb.Provider.AddEnv(constants.DatabaseNameSpace, env.DatabaseNameSpace)
	memorydb.Provider.AddEnv(constants.DatabaseUsername, env.DatabaseUsername)
	memorydb.Provider.AddEnv(constants.DatabasePassword, env.DatabasePassword)
	memorydb.Provider.AddEnv(constants.Port, env.Port)
	memorydb.Provider.AddEnv(constants.EncryptionKey, env.EncryptionKey)
	memorydb.Provider.AddEnv(constants.JwtType, env.JwtType)
	memorydb.Provider.AddEnv(constants.JwtSecret, env.JwtSecret)
	memorydb.Provider.AddEnv(constants.JwtPrivateKey, env.JwtPrivateKey)
	memorydb.Provider.AddEnv(constants.JwtPublicKey, env.JwtPublicKey)
	memorydb.Provider.AddEnv(constants.ClientID, env.ClientID)

	return env
}
