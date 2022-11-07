package pebbledb

import (
	"github.com/cockroachdb/pebble"
	log "github.com/sirupsen/logrus"
)

func (p *provider) SetSession(key string, value string) (string, error) {
	err := p.db.Set([]byte(key), []byte(value), pebble.NoSync)
	if err != nil {
		log.Debug("Error setting session in pebbledb : ", err)
	}

	return key, nil
}

func (p *provider) GetSession(key string) (string, error) {
	value, closer, err := p.db.Get([]byte(key))
	if err != nil {
		log.Debug("Error getting session by key from pebbledb : ", err)
	}
	closer.Close()

	return string(value), nil
}

func (p *provider) DeleteSession(key string) error {
	err := p.db.Delete([]byte(key), pebble.NoSync)
	if err != nil {
		log.Debug("Error deleting session in pebbledb : ", err)
	}

	return nil
}
