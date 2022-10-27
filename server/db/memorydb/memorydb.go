package memorydb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/memorydb/providers/genjidb"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMemoryDB() error {
	var err error

	Provider, err = genjidb.NewGenjiDbProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing genjidb")
		return err
	} else {
		logrus.Info("Genjidb successfully initialized")
	}

	return nil
}
