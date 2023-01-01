package db

import (
	"gategaurdian/server/constants"
	"gategaurdian/server/database/db/providers"
	"gategaurdian/server/database/db/providers/sql"
	"gategaurdian/server/database/memorystore"

	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

// InitMainDb is used to initialize the main database used by the application
func InitMainDb() error {
	var err error

	dbType := memorystore.RequiredEnvStoreObj.GetRequiredEnv().DatabaseType

	isSql := dbType != constants.DbTypeMongoDb && dbType != constants.DbTypeCassandraDb && dbType != constants.DbTypeScyllaDb

	if isSql {
		log.Info("Initializing SQl driver for : ", dbType)
		Provider, err = sql.NewProvider()
		if err != nil {
			log.Fatal("Failed to initialize SQL driver: ", err)
			return err
		}
	}

	return nil
}
