package maindb

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/providers/sql"
	"github.com/ArkamFahry/GateGuardian/server/env"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMainDB() error {
	var err error

	dbType, _ := env.GetEnvByKey(constants.DatabaseType)
	isSql := dbType != constants.DbTypeMongodb

	if isSql {
		Provider, err = sql.NewSqlProvider()
		if err != nil {
			logrus.Fatal("Failed Initializing sqlite: ", err)
			return err
		} else {
			logrus.Info("Sqlite successfully initialized")
		}
	}

	return nil
}
