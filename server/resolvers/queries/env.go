package queries

import (
	"context"
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
)

func EnvResolver(ctx context.Context) (*model.Env, error) {
	res := &model.Env{}

	dbUrl, _ := env.GetEnvByKey(constants.DatabaseURL)
	res.DatabaseURL = &dbUrl

	dbName, _ := env.GetEnvByKey(constants.DatabaseName)
	res.DatabaseName = &dbName

	dbNameSpace, _ := env.GetEnvByKey(constants.DatabaseNameSpace)
	res.DatabaseNamespace = &dbNameSpace

	dbUserName, _ := env.GetEnvByKey(constants.DatabaseUsername)
	res.DatabaseUsername = &dbUserName

	port, _ := env.GetEnvByKey(constants.Port)
	res.Port = &port

	jwtType, _ := env.GetEnvByKey(constants.JwtType)
	res.JwtType = &jwtType

	jwtSecret, _ := env.GetEnvByKey(constants.JwtSecret)
	res.JwtSecret = &jwtSecret

	jwtPrivateKey, _ := env.GetEnvByKey(constants.JwtPrivateKey)
	res.JwtPrivateKey = &jwtPrivateKey

	jwtPublicKey, _ := env.GetEnvByKey(constants.JwtPublicKey)
	res.JwtPublicKey = &jwtPublicKey

	clientId, _ := env.GetEnvByKey(constants.ClientID)
	res.ClientID = &clientId

	roles, _ := env.GetEnvByKey(constants.Roles)
	res.Roles = strings.Split(roles, ",")

	defaultRoles, _ := env.GetEnvByKey(constants.DefaultRoles)
	res.DefaultRoles = strings.Split(defaultRoles, ",")

	return res, nil
}
