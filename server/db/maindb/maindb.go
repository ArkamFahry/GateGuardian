package maindb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/providers/surrealdb"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMainDB() error {
	var err error

	Provider, err = surrealdb.NewSurrealDbProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing surrealdb: ", err)
		return err
	} else {
		logrus.Info("Surrealdb successfully initialized")
	}

	return nil
}
