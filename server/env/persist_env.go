package env

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/memorydb"
)

func PersistEnv(env Env) {
	EncryptionKey := crypto.EncryptB64(env.EncryptionKey)
	memorydb.Provider.AddEnv(constants.EncryptionKey, EncryptionKey)

	AddEnv(constants.DatabaseType, env.DatabaseType)
	AddEnv(constants.DatabaseURL, env.DatabaseURL)
	AddEnv(constants.DatabaseName, env.DatabaseName)
	AddEnv(constants.DatabaseNameSpace, env.DatabaseNameSpace)
	AddEnv(constants.DatabaseUsername, env.DatabaseUsername)
	AddEnv(constants.DatabasePassword, env.DatabasePassword)
	AddEnv(constants.Port, env.Port)
	AddEnv(constants.JwtType, env.JwtType)
	AddEnv(constants.JwtSecret, env.JwtSecret)
	AddEnv(constants.JwtPrivateKey, env.JwtPrivateKey)
	AddEnv(constants.JwtPublicKey, env.JwtPublicKey)
	AddEnv(constants.ClientID, env.ClientID)
	AddEnv(constants.DefaultRoles, env.DefaultRoles)
	AddEnv(constants.Roles, env.Roles)
}

func AddEnv(key string, data string) (string, error) {
	env, err := crypto.EncryptAES(data)
	if err != nil {
		return env, err
	}
	env, err = memorydb.Provider.AddEnv(key, env)

	return env, err
}

func GetEnvByKey(key string) (string, error) {
	env, err := memorydb.Provider.GetEnvByKey(key)
	if err != nil {
		return env, err
	}
	env, err = crypto.DecryptAES(env)

	return env, err
}

func UpdateEnv(key string, data string) (string, error) {
	env, err := crypto.EncryptAES(data)
	if err != nil {
		return env, err
	}
	env, err = memorydb.Provider.UpdateEnv(key, env)

	return env, err
}
