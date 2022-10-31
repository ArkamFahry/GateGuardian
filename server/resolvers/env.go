package resolvers

import (
	"context"
	"strings"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/crypto"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
	"github.com/ArkamFahry/GateGuardian/server/memorydb"
)

func EnvResolver(ctx context.Context) (*model.Env, error) {
	res := &model.Env{}

	dbUrl, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseURL)
	dbUrl, _ = crypto.DecryptAES(dbUrl)
	res.DatabaseURL = &dbUrl

	dbName, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseName)
	dbName, _ = crypto.DecryptAES(dbName)
	res.DatabaseName = &dbName

	dbNameSpace, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseNameSpace)
	dbNameSpace, _ = crypto.DecryptAES(dbNameSpace)
	res.DatabaseNamespace = &dbNameSpace

	dbUserName, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseUsername)
	dbUserName, _ = crypto.DecryptAES(dbUserName)
	res.DatabaseUsername = &dbUserName

	port, _ := memorydb.Provider.GetEnvByKey(constants.Port)
	port, _ = crypto.DecryptAES(port)
	res.Port = &port

	jwtType, _ := memorydb.Provider.GetEnvByKey(constants.JwtType)
	jwtType, _ = crypto.DecryptAES(jwtType)
	res.JwtType = &jwtType

	jwtSecret, _ := memorydb.Provider.GetEnvByKey(constants.JwtSecret)
	jwtSecret, _ = crypto.DecryptAES(jwtSecret)
	res.JwtSecret = &jwtSecret

	jwtPrivateKey, _ := memorydb.Provider.GetEnvByKey(constants.JwtPrivateKey)
	jwtPrivateKey, _ = crypto.DecryptAES(jwtPrivateKey)
	res.JwtPrivateKey = &jwtPrivateKey

	jwtPublicKey, _ := memorydb.Provider.GetEnvByKey(constants.JwtPublicKey)
	jwtPublicKey, _ = crypto.DecryptAES(jwtPublicKey)
	res.JwtPublicKey = &jwtPublicKey

	clientId, _ := memorydb.Provider.GetEnvByKey(constants.ClientID)
	clientId, _ = crypto.DecryptAES(clientId)
	res.ClientID = &clientId

	roles, _ := memorydb.Provider.GetEnvByKey(constants.Roles)
	roles, _ = crypto.DecryptAES(roles)
	res.Roles = strings.Split(roles, ",")

	defaultRoles, _ := memorydb.Provider.GetEnvByKey(constants.DefaultRoles)
	defaultRoles, _ = crypto.DecryptAES(defaultRoles)
	res.DefaultRoles = strings.Split(defaultRoles, ",")

	return res, nil
}
