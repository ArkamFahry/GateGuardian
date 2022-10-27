package surrealdb

import (
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/sirupsen/logrus"
	"github.com/surrealdb/surrealdb.go"
)

type provider struct {
	maindb *surrealdb.DB
}

func NewSurrealDbProvider() (*provider, error) {
	dbUrl := env.EnvGet().DatabaseURL
	surrealdb, err := surrealdb.New(dbUrl)
	if err != nil {
		logrus.Fatal("Failed to create a connection to surrealdb: ", err)
		return nil, err
	}

	dbUserName := env.EnvGet().DatabaseUsername
	dbPassword := env.EnvGet().DatabasePassword
	_, err = surrealdb.Signin(map[string]any{
		"user": dbUserName,
		"pass": dbPassword,
	})
	if err != nil {
		logrus.Fatal("Failed to sign in to surrealdb with username and password: ", err)
	}

	dbName := env.EnvGet().DatabaseName
	dbNameSpace := env.EnvGet().DatabaseNameSpace
	_, err = surrealdb.Use(dbNameSpace, dbName)
	if err != nil {
		logrus.Fatal("Failed to select surrealdb database or namespace: ", err)
	}

	return &provider{
		maindb: surrealdb,
	}, err
}
