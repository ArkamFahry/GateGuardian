package env

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
)

func PersistEnvToCache() {
	memorydb.Provider.AddEnv(constants.DatabaseURL, EnvGet().DatabaseURL)
	memorydb.Provider.AddEnv(constants.DatabaseName, EnvGet().DatabaseName)
	memorydb.Provider.AddEnv(constants.DatabaseNameSpace, EnvGet().DatabaseNameSpace)
	memorydb.Provider.AddEnv(constants.DatabaseUsername, EnvGet().DatabaseUsername)
	memorydb.Provider.AddEnv(constants.DatabasePassword, EnvGet().DatabasePassword)
	memorydb.Provider.AddEnv(constants.Port, EnvGet().Port)
	memorydb.Provider.AddEnv(constants.EncryptionKey, EnvGet().EncryptionKey)
	memorydb.Provider.AddEnv(constants.JwtType, EnvGet().JwtType)
	memorydb.Provider.AddEnv(constants.JwtSecret, EnvGet().JwtSecret)
	memorydb.Provider.AddEnv(constants.ClientID, EnvGet().ClientID)
}
