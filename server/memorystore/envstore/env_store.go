package envstore

import (
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore/providers"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore/providers/pebbledb"
	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitEnvStore() error {
	var err error

	Provider, err = pebbledb.NewPebbleDbProvider()
	if err != nil {
		log.Debug("Failed to initialize env store : ", err)
	}

	return err
}
