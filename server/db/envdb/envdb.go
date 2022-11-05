package envdb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/envdb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/envdb/providers/memdb"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitEnvDB() error {
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
