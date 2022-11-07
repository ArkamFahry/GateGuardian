package pebbledb

import (
	"github.com/cockroachdb/pebble"
	log "github.com/sirupsen/logrus"
)

func (p *provider) SetEnv(key string, value string) (string, error) {
	err := p.db.Set([]byte(key), []byte(value), pebble.NoSync)
	if err != nil {
		log.Debug("Error setting env in pebbledb : ", err)
	}

	return key, nil
}

func (p *provider) GetEnv(key string) (string, error) {
	value, closer, err := p.db.Get([]byte(key))
	if err != nil {
		log.Debug("Error getting env by key from pebbledb : ", err)
	}
	closer.Close()

	return string(value), nil
}

func (p *provider) DeleteEnv(key string) error {
	err := p.db.Delete([]byte(key), pebble.NoSync)
	if err != nil {
		log.Debug("Error deleting env in pebbledb : ", err)
	}

	return nil
}
