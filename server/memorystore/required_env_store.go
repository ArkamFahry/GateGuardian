package memorystore

import (
	"os"
	"strings"
	"sync"

	"github.com/ArkamFahry/GateGuardian/server/cli"
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// RequiredEnv holds information about required environment variables
type RequiredEnv struct {
	EnvPath           string `json:"ENV_PATH"`
	DatabaseURL       string `json:"DATABASE_URL"`
	DatabaseType      string `json:"DATABASE_TYPE"`
	DatabaseName      string `json:"DATABASE_NAME"`
	DatabaseNameSpace string `json:"DATABASE_NAMESPACE"`
	DatabaseHost      string `json:"DATABASE_HOST"`
	DatabasePort      string `json:"DATABASE_PORT"`
	DatabaseUsername  string `json:"DATABASE_USERNAME"`
	DatabasePassword  string `json:"DATABASE_PASSWORD"`
	DatabaseCert      string `json:"DATABASE_CERT"`
	DatabaseCertKey   string `json:"DATABASE_CERT_KEY"`
	DatabaseCACert    string `json:"DATABASE_CA_CERT"`
}

// RequiredEnvStore is a simple in-memory store for sessions.
type RequiredEnvStore struct {
	mutex       sync.Mutex
	requiredEnv RequiredEnv
}

// GetRequiredEnv to get required environment variables
func (r *RequiredEnvStore) GetRequiredEnv() RequiredEnv {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.requiredEnv
}

// SetRequiredEnv to set required environment variables
func (r *RequiredEnvStore) SetRequiredEnv(requiredEnv RequiredEnv) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.requiredEnv = requiredEnv
}

var RequiredEnvStoreObj *RequiredEnvStore

// InitRequiredEnv to initialize EnvData and throw error if required env are not present
// This includes env that only configurable via env vars and not the ui
func InitRequiredEnv() error {
	envPath := os.Getenv(constants.EnvKeyEnvPath)

	if envPath == "" {
		if envPath == "" {
			envPath = `.env`
		}
	}

	if cli.ARG_ENV_FILE != nil && *cli.ARG_ENV_FILE != "" {
		envPath = *cli.ARG_ENV_FILE
	}
	logrus.Info("env path: ", envPath)

	err := godotenv.Load(envPath)
	if err != nil {
		logrus.Infof("using OS env instead of %s file", envPath)
	}

	dbURL := os.Getenv(constants.EnvKeyDatabaseURL)
	dbName := os.Getenv(constants.EnvKeyDatabaseName)
	dbNameSpace := os.Getenv(constants.EnvKeyDatabaseNameSpace)
	dbPort := os.Getenv(constants.EnvKeyDatabasePort)
	dbHost := os.Getenv(constants.EnvKeyDatabaseHost)
	dbUsername := os.Getenv(constants.EnvKeyDatabaseUsername)
	dbPassword := os.Getenv(constants.EnvKeyDatabasePassword)
	dbCert := os.Getenv(constants.EnvKeyDatabaseCert)
	dbCertKey := os.Getenv(constants.EnvKeyDatabaseCertKey)
	dbCACert := os.Getenv(constants.EnvKeyDatabaseCACert)

	// set default db name for non sql dbs
	if dbName == "" {
		dbName = "gate_guardian"
	}

	if strings.TrimSpace(dbURL) == "" {
		if cli.ARG_DB_URL != nil && *cli.ARG_DB_URL != "" {
			dbURL = strings.TrimSpace(*cli.ARG_DB_URL)
		}
	}

	if dbName == "" {
		if dbName == "" {
			dbName = "gate_guardian"
		}
	}

	requiredEnv := RequiredEnv{
		EnvPath:           envPath,
		DatabaseURL:       dbURL,
		DatabaseName:      dbName,
		DatabaseNameSpace: dbNameSpace,
		DatabaseHost:      dbHost,
		DatabasePort:      dbPort,
		DatabaseUsername:  dbUsername,
		DatabasePassword:  dbPassword,
		DatabaseCert:      dbCert,
		DatabaseCertKey:   dbCertKey,
		DatabaseCACert:    dbCACert,
	}

	RequiredEnvStoreObj = &RequiredEnvStore{
		requiredEnv: requiredEnv,
	}

	return nil
}
