package maindb

import (
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/providers"
	"github.com/ArkamFahry/GateGuardian/server/db/maindb/providers/sqlite"
	"github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitMainDB() error {
	var err error

	Provider, err = sqlite.NewSqliteProvider()
	if err != nil {
		logrus.Fatal("Failed Initializing sqlite: ", err)
		return err
	} else {
		logrus.Info("Sqlite successfully initialized")
	}

	return nil
}
