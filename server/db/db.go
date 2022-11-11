package db

import (
	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/db/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/providers/sqlite"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitDB() error {
	var err error

	dbType, _ := envstore.Provider.GetEnv(constants.DbType)

	isSqlite := dbType == constants.DbTypeSqlite

	if isSqlite {
		log.Info("Initialized db : ", dbType)
		Provider, err = sqlite.NewProvider()
		if err != nil {
			log.Fatal("Failed to initialize SQL driver: ", err)
			return err
		}
	}

	return nil
}
