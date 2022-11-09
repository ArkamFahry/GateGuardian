package env

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
)

func PersistEnv(envs Envs) {
	envstore.Provider.SetEnv(constants.PORT, envs.PORT)
	envstore.Provider.SetEnv(constants.DB_TYPE, envs.DB_TYPE)
	envstore.Provider.SetEnv(constants.DB_URL, envs.DB_URL)
}
