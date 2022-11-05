package envdb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/envdb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/envdb/providers/pebbledb"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitEnvDB() error {
	var err error

	Provider, err = pebbledb.NewPebbleDbProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing pebbledb env store: ", err)
		return err
	} else {
		logrus.Info("PebbleDb env store successfully initialized")
	}

	return nil
}
