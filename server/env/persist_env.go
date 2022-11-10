package env

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
)

func PersistEnv(envs Envs) {
	if envs.PORT == "" {
		envs.PORT = "8080"
	}
	envstore.Provider.SetEnv(constants.PORT, envs.PORT)

	if envs.DB_TYPE == "" {
		envs.DB_TYPE = "sqlite"
	}
	envstore.Provider.SetEnv(constants.DB_TYPE, envs.DB_TYPE)

	if envs.DB_URL == "" {
		envs.DB_URL = "../data.db"
	}
	envstore.Provider.SetEnv(constants.DB_URL, envs.DB_URL)
}
