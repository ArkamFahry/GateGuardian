package surrealdb

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/memorydb"
	"github.com/sirupsen/logrus"
	"github.com/surrealdb/surrealdb.go"
)

type provider struct {
	db *surrealdb.DB
}

func NewSurrealDbProvider() (*provider, error) {
	dbUrl, err := memorydb.Provider.GetEnvByKey(constants.DatabaseURL)
	if err != nil {
		logrus.Fatal("Couldn't load dbUrl from cache: ", err)
	}

	surrealdb, err := surrealdb.New(dbUrl)
	if err != nil {
		logrus.Fatal("Failed to create a connection to surrealdb: ", err)
		return nil, err
	}

	dbUserName, err := memorydb.Provider.GetEnvByKey(constants.DatabaseUsername)
	if err != nil {
		logrus.Fatal("Couldn't load dbUserName from cache: ", err)
	}
	dbPassword, _ := memorydb.Provider.GetEnvByKey(constants.DatabasePassword)
	if err != nil {
		logrus.Fatal("Couldn't load dbPassword from cache: ", err)
	}

	_, err = surrealdb.Signin(map[string]any{
		"user": dbUserName,
		"pass": dbPassword,
	})
	if err != nil {
		logrus.Fatal("Failed to sign in to surrealdb with username and password: ", err)
	}

	dbName, err := memorydb.Provider.GetEnvByKey(constants.DatabaseName)
	if err != nil {
		logrus.Fatal("Couldn't load dbName from cache: ", err)
	}
	dbNameSpace, err := memorydb.Provider.GetEnvByKey(constants.DatabaseNameSpace)
	if err != nil {
		logrus.Fatal("Couldn't load dbNameSpace from cache: ", err)
	}

	_, err = surrealdb.Use(dbNameSpace, dbName)
	if err != nil {
		logrus.Fatal("Failed to select surrealdb database or namespace: ", err)
	}

	return &provider{
		db: surrealdb,
	}, err
}
