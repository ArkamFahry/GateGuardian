package resolvers

import (
	"context"
	"fmt"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb"
	"github.com/ArkamFahry/GateGuardian/server/graph/model"
)

func EnvResolver(ctx context.Context) (*model.Env, error) {
	res := &model.Env{}

	dbUrl, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseURL)
	res.DatabaseURL = &dbUrl
	dbName, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseName)
	res.DatabaseName = &dbName
	dbNameSpace, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseNameSpace)
	res.DatabaseNamespace = &dbNameSpace
	dbUserName, _ := memorydb.Provider.GetEnvByKey(constants.DatabaseUsername)
	res.DatabaseUsername = &dbUserName
	port, _ := memorydb.Provider.GetEnvByKey(constants.Port)
	res.Port = &port
	jwtType, _ := memorydb.Provider.GetEnvByKey(constants.JwtType)
	res.JwtType = &jwtType
	jwtSecret, _ := memorydb.Provider.GetEnvByKey(constants.JwtSecret)
	res.JwtSecret = &jwtSecret
	clientId, _ := memorydb.Provider.GetEnvByKey(constants.ClientID)
	res.ClientID = &clientId

	fmt.Println(memorydb.Provider.ListEnv())

	return res, nil
}
