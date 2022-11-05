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
		logrus.Fatal("Failed Initializing pebbledb memory store: ", err)
		return err
	} else {
		logrus.Info("PebbleDb memory store successfully initialized")
	}

	return nil
}
