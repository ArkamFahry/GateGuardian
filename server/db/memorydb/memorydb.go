package memorydb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/providers/pebbledb"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMemoryDB() error {
	var err error

	Provider, err = pebbledb.NewPebbleDbProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing pebbledb: ", err)
		return err
	} else {
		logrus.Info("PebbleDb successfully initialized")
	}

	return nil
}
