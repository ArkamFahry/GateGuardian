package memorydb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/providers/genjidb"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/providers/sqlite"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMemoryDB() error {
	var err error

	Provider, err = genjidb.NewGenjiDbProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing genjidb: ", err)
		return err
	} else {
		logrus.Info("GenjiDb successfully initialized")
	}

	Provider, err = sqlite.NewSqliteProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing sqlite: ", err)
		return err
	} else {
		logrus.Info("Sqlite successfully initialized")
	}

	return nil
}
