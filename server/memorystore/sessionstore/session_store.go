package sessionstore

import (
	"github.com/ArkamFahry/GateGuardian/server/memorystore/sessionstore/providers"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/sessionstore/providers/pebbledb"
	log "github.com/sirupsen/logrus"
)

var Provider providers.Provider

func InitSessionStore() error {
	var err error

	Provider, err = pebbledb.NewProvider()
	if err != nil {
		log.Debug("Failed to initialize session store : ", err)
	}

	return err
}
