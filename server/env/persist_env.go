package env

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/db/envdb"
)

func PersistEnv(env Env) {
	EncryptionKey := crypto.EncryptB64(env.EncryptionKey)
	envdb.Provider.AddEnv(constants.EncryptionKey, EncryptionKey)

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
	AddEnv(constants.JwtRoleClaim, env.JwtRoleClaim)
	AddEnv(constants.ClientID, env.ClientID)
	AddEnv(constants.AccessTokenExpiryTime, env.AccessTokenExpiryTime)
	AddEnv(constants.CustomAccessTokenScript, env.CustomAccessTokenScript)
	AddEnv(constants.DefaultRoles, env.DefaultRoles)
	AddEnv(constants.Roles, env.Roles)
	AddEnv(constants.AppURL, env.AppURL)
	AddEnv(constants.GateGuardianURL, env.GateGuardianURL)
}

func AddEnv(key string, data string) (string, error) {
	env, err := crypto.EncryptAES(data)
	if err != nil {
		return env, err
	}
	env, err = envdb.Provider.AddEnv(key, env)

	return env, err
}

func GetEnvByKey(key string) (string, error) {
	env, err := envdb.Provider.GetEnvByKey(key)
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
	env, err = envdb.Provider.UpdateEnv(key, env)

	return env, err
}
