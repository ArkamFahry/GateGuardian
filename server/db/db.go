package db

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/providers/genjidb"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMainDB() error {
	var err error

	DbType, _ := env.GetEnvByKey(constants.DatabaseType)

	isGenjiDb := DbType == constants.DbTypeGenjiDb
	// isSurrealDb := DbType == constants.DbTypeSurrealDb

	if isGenjiDb {
		Provider, err = genjidb.NewGenjiDbProvider()
		if err != nil {
			logrus.Fatal("Failed Initializing genjidb: ", err)
			return err
		} else {
			logrus.Info("Genjidb successfully initialized")
		}
	}

	// if isSurrealDb {
	// 	Provider, err = surrealdb.NewSurrealDbProvider()
	// 	if err != nil {
	// 		logrus.Fatal("Failed Initializing surrealdb: ", err)
	// 		return err
	// 	} else {
	// 		logrus.Info("Surrealdb successfully initialized")
	// 	}
	// }

	return nil
}
