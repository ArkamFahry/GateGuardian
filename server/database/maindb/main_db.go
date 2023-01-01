package maindb

import (
	"gategaurdian/server/constants"
	"gategaurdian/server/database/maindb/providers"
	"gategaurdian/server/database/maindb/providers/sql"
	"gategaurdian/server/database/memorydb"

	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

// InitMainDb is used to initialize the main database used by the application
func InitMainDb() error {
	var err error

	dbType := memorydb.RequiredEnvStoreObj.GetRequiredEnv().DatabaseType

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
