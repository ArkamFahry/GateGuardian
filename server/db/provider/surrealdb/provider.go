package surrealdb

import (
	"github.com/ArkamFahry/GateGuardian/server/memorystore"
	"github.com/surrealdb/surrealdb.go"
)

type provider struct {
	db *surrealdb.DB
}

func NewSurrealDbProvider() (*provider, error) {
	dbUrl := memorystore.RequiredEnvStoreObj.GetRequiredEnv().DatabaseURL
	surrealdb, err := surrealdb.New(dbUrl)
	if err != nil {
		return nil, err
	}

	dbUserName := memorystore.RequiredEnvStoreObj.GetRequiredEnv().DatabaseUsername
	dbPassword := memorystore.RequiredEnvStoreObj.GetRequiredEnv().DatabasePassword
	surrealdb.Signin(map[string]any{
		"user": dbUserName,
		"pass": dbPassword,
	})

	dbName := memorystore.RequiredEnvStoreObj.GetRequiredEnv().DatabaseName
	dbNameSpace := memorystore.RequiredEnvStoreObj.GetRequiredEnv().DatabaseNameSpace
	surrealdb.Use(dbNameSpace, dbName)

	return &provider{
		db: surrealdb,
	}, nil
}
