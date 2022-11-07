package pebbledb

import (
	"github.com/cockroachdb/pebble"
	log "github.com/sirupsen/logrus"
)

func (p *provider) AddEnv(key string, data string) (string, error) {
	err := p.db.Set([]byte(key), []byte(data), pebble.Sync)
	if err != nil {
		log.Debug("pebbledb error can't set value")
	}

	return key, nil
}

func (p *provider) UpdateEnv(key string, data string) (string, error) {
	err := p.db.Set([]byte(key), []byte(data), pebble.Sync)
	if err != nil {
		log.Debug("pebbledb error can't update value")
	}

	return key, nil
}

func (p *provider) DeleteEnv(key string) error {
	err := p.db.Delete([]byte(key), pebble.Sync)
	if err != nil {
		log.Debug("pebbledb error can't update value")
	}

	return nil
}

func (p *provider) GetEnvByKey(key string) (string, error) {
	env, closer, err := p.db.Get([]byte(key))
	if err != nil {
		log.Debug("pebbledb error can't get value")
	}

	closer.Close()

	return string(env), nil
}
