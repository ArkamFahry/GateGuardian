package pebbledb

import (
	"github.com/cockroachdb/pebble"
	"github.com/sirupsen/logrus"
)

func (p *provider) SetSession(key string, value string) (string, error) {
	err := p.db.Set([]byte(key), []byte(value), pebble.Sync)
	if err != nil {
		logrus.Info("pebbledb error can't set value")
	}

	return key, nil
}

func (p *provider) GetSession(key string) (string, error) {
	value, closer, err := p.db.Get([]byte(key))
	if err != nil {
		logrus.Info("pebbledb error can't get value")
	}

	closer.Close()

	return string(value), nil
}

func (p *provider) UpdateSession(key string, value string) (string, error) {
	err := p.db.Set([]byte(key), []byte(value), pebble.Sync)
	if err != nil {
		logrus.Info("pebbledb error can't update value")
	}

	return key, nil
}
func (p *provider) DeleteSession(key string) error {
	err := p.db.Delete([]byte(key), pebble.Sync)
	if err != nil {
		logrus.Info("pebbledb error can't update value")
	}

	return nil
}
