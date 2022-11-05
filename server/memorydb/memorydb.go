package memorydb

import (
	"github.com/ArkamFahry/GateGuardian/server/memorydb/providers"
	"github.com/ArkamFahry/GateGuardian/server/memorydb/providers/memdb"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMemoryDB() error {
	var err error

	Provider, err = memdb.NewMemDbProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing genjidb: ", err)
		return err
	} else {
		logrus.Info("GenjiDb successfully initialized")
	}

	return nil
}
