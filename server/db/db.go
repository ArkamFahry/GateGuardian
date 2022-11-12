package db

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/providers/sql"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitDB() error {
	var err error

	dbType, _ := envstore.Provider.GetEnv(constants.DbType)

	isSql := dbType == constants.DbTypeSqlite || dbType == constants.DbTypeMysql || dbType == constants.DbTypePostgresql || dbType == constants.DbTypeSqlserver

	if isSql {
		log.Info("Initialized db : ", dbType)
		Provider, err = sql.NewProvider()
		if err != nil {
			log.Fatal("Failed to initialize SQL driver: ", err)
			return err
		}
	}

	return nil
}
