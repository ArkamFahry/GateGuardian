package db

import (
	"github.com/ArkamFahry/GateGuardian/server/db/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/providers/surrealdb"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitDB() error {
	var err error

	Provider, err = surrealdb.NewSurrealDbProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing surrealdb")
		return err
	} else {
		logrus.Info("Surrealdb successfully initialized")
	}

	return nil
}
